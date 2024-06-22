package service

import (
	"context"
	"strconv"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	ptr "github.com/teran/go-ptr"

	"github.com/runityru/cephctl/ceph"
	"github.com/runityru/cephctl/differ"
	"github.com/runityru/cephctl/models"
	clusterHeath "github.com/runityru/cephctl/service/cluster_health"
)

func init() {
	log.SetLevel(log.TraceLevel)
}

func (s *serviceTestSuite) TestApplyCephConfig() {
	currentConfig := models.CephConfig{
		"osd": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "old_value",
		},
	}
	newConfig := models.CephConfig{
		"mon": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "value",
		},
	}
	result := []models.CephConfigDifference{
		{
			Kind:    models.CephConfigDifferenceKindAdd,
			Section: "mon",
			Key:     "test_key",
			Value:   ptr.String("value"),
		},
		{
			Kind:     models.CephConfigDifferenceKindChange,
			Section:  "osd.3",
			Key:      "test_key",
			OldValue: ptr.String("old_value"),
			Value:    ptr.String("value"),
		},
		{
			Kind:    models.CephConfigDifferenceKindRemove,
			Section: "osd",
			Key:     "test_key",
		},
	}

	cephDumpConfig := s.cephMock.On("DumpConfig").Return(currentConfig, nil).Once()

	s.differMock.On("DiffCephConfig", currentConfig, newConfig).Return(result, nil).Once()

	s.cephMock.On("RemoveCephConfigOption", "osd", "test_key").Return(nil).NotBefore(cephDumpConfig).Once()
	s.cephMock.On("ApplyCephConfigOption", "mon", "test_key", "value").Return(nil).NotBefore(cephDumpConfig).Once()
	s.cephMock.On("ApplyCephConfigOption", "osd.3", "test_key", "value").Return(nil).NotBefore(cephDumpConfig).Once()

	err := s.svc.ApplyCephConfig(s.ctx, newConfig)
	s.Require().NoError(err)
}

func (s *serviceTestSuite) TestApplyCephOSDConfig() {
	newCfg := models.CephOSDConfig{
		AllowCrimson:           true,
		BackfillfullRatio:      0.95,
		FullRatio:              0.98,
		NearfullRatio:          0.9,
		RequireMinCompatClient: "reef",
	}

	call1 := s.cephMock.On("ClusterReport").Return(models.ClusterReport{}, nil).Once()
	call2 := s.differMock.On("DiffCephOSDConfig", models.CephOSDConfig{}, newCfg).Return([]models.CephOSDConfigDifference{
		{
			Key:      "blah",
			OldValue: "932",
			Value:    "123",
		},
		{
			Key:      "some_key",
			OldValue: "123",
			Value:    "456",
		},
	}, nil).NotBefore(call1).Once()
	call3 := s.cephMock.On("ApplyCephOSDConfigOption", "blah", "123").Return(nil).NotBefore(call2).Once()
	s.cephMock.On("ApplyCephOSDConfigOption", "some_key", "456").Return(nil).NotBefore(call3).Once()

	err := s.svc.ApplyCephOSDConfig(s.ctx, newCfg)
	s.Require().NoError(err)
}

func (s *serviceTestSuite) TestCheckClusterHealth() {
	s.cephMock.On("ClusterReport").Return(models.ClusterReport{
		HealthStatus:    models.ClusterStatusHealthOK,
		Checks:          []models.ClusterStatusCheck{},
		MutedChecks:     []models.ClusterStatusMutedCheck{},
		NumMons:         5,
		NumMonsInQuorum: 5,
		NumOSDs:         15,
		NumOSDsIn:       15,
		NumOSDsUp:       15,
		NumOSDsByRelease: map[string]uint16{
			"reef": 15,
		},
		NumOSDsByVersion: map[string]uint16{
			"18.2.2": 15,
		},
		NumOSDsByDeviceType: map[string]uint16{
			"ssd": 15,
		},
		TotalOSDCapacityKB: uint64(22_321_704_960),
		TotalOSDUsedDataKB: uint64(10_986_978_208),
		TotalOSDUsedMetaKB: uint64(512_967_627),
		TotalOSDUsedOMAPKB: uint64(5_822_580),
		NumPools:           14,
		NumPGs:             330,
		NumPGsByState: map[string]uint32{
			"active":        330,
			"backfill_wait": 50,
			"backfilling":   2,
			"clean":         278,
			"remapped":      52,
		},
	}, nil).Once()

	s.cephMock.On("ListDevices").Return([]models.Device{
		{
			ID:        "testdevice",
			Daemons:   []string{"osd.0"},
			WearLevel: 0.000001,
		},
		{
			ID:        "testdevice2",
			Daemons:   []string{"osd.0"},
			WearLevel: 0.510001,
		},
	}, nil)

	chi, err := s.svc.CheckClusterHealth(s.ctx, []clusterHeath.ClusterHealthCheck{
		func(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
			return models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
				CurrentValue:       "HEALTH_OK",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			}, nil
		},
	})
	s.Require().NoError(err)
	s.Require().Equal([]models.ClusterHealthIndicator{
		{
			Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
			CurrentValue:       "HEALTH_OK",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
	}, chi)
}

func (s *serviceTestSuite) TestDiffCephConfig() {
	currentConfig := models.CephConfig{
		"osd": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "old_value",
		},
	}
	newConfig := models.CephConfig{
		"mon": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "value",
		},
	}
	result := []models.CephConfigDifference{
		{
			Kind:    models.CephConfigDifferenceKindAdd,
			Section: "mon",
			Key:     "test_key",
			Value:   ptr.String("value"),
		},
		{
			Kind:     models.CephConfigDifferenceKindChange,
			Section:  "osd.3",
			Key:      "test_key",
			OldValue: ptr.String("old_value"),
			Value:    ptr.String("value"),
		},
		{
			Kind:    models.CephConfigDifferenceKindRemove,
			Section: "osd",
			Key:     "test_key",
		},
	}

	cephDumpConfig := s.cephMock.
		On("DumpConfig").Return(currentConfig, nil).Once()
	s.differMock.
		On("DiffCephConfig", currentConfig, newConfig).Return(result, nil).NotBefore(cephDumpConfig).Once()

	diff, err := s.svc.DiffCephConfig(s.ctx, newConfig)
	s.Require().NoError(err)
	s.Require().ElementsMatch(result, diff)
}

func (s *serviceTestSuite) TestDiffCephOSDConfig() {
	src := models.CephOSDConfig{
		AllowCrimson:           false,
		NearfullRatio:          0.85,
		BackfillfullRatio:      0.90,
		FullRatio:              0.95,
		RequireMinCompatClient: "luminous",
	}

	dst := models.CephOSDConfig{
		AllowCrimson:           true,
		NearfullRatio:          0.89,
		BackfillfullRatio:      0.92,
		FullRatio:              0.97,
		RequireMinCompatClient: "reef",
	}

	s.cephMock.On("ClusterReport").Return(models.ClusterReport{
		AllowCrimson:           false,
		NearfullRatio:          0.85,
		BackfillfullRatio:      0.90,
		FullRatio:              0.95,
		RequireMinCompatClient: "luminous",
	}, nil).Once()

	s.differMock.On("DiffCephOSDConfig", src, dst).Return([]models.CephOSDConfigDifference{
		{
			Key:      "AllowCrimson",
			OldValue: "false",
			Value:    "true",
		},
		{
			Key:      "NearfullRatio",
			OldValue: strconv.FormatFloat(0.85, 'f', 2, 32),
			Value:    strconv.FormatFloat(0.89, 'f', 2, 32),
		},
		{
			Key:      "BackfillfullRatio",
			OldValue: strconv.FormatFloat(0.90, 'f', 2, 32),
			Value:    strconv.FormatFloat(0.92, 'f', 2, 32),
		},
		{
			Key:      "FullRatio",
			OldValue: strconv.FormatFloat(0.95, 'f', 2, 32),
			Value:    strconv.FormatFloat(0.97, 'f', 2, 32),
		},
		{
			Key:      "RequireMinCompatClient",
			OldValue: "luminous",
			Value:    "reef",
		},
	}, nil).Once()

	diff, err := s.svc.DiffCephOSDConfig(s.ctx, dst)
	s.Require().NoError(err)
	s.Require().ElementsMatch([]models.CephOSDConfigDifference{
		{
			Key:      "AllowCrimson",
			OldValue: "false",
			Value:    "true",
		},
		{
			Key:      "NearfullRatio",
			OldValue: strconv.FormatFloat(0.85, 'f', 2, 32),
			Value:    strconv.FormatFloat(0.89, 'f', 2, 32),
		},
		{
			Key:      "BackfillfullRatio",
			OldValue: strconv.FormatFloat(0.90, 'f', 2, 32),
			Value:    strconv.FormatFloat(0.92, 'f', 2, 32),
		},
		{
			Key:      "FullRatio",
			OldValue: strconv.FormatFloat(0.95, 'f', 2, 32),
			Value:    strconv.FormatFloat(0.97, 'f', 2, 32),
		},
		{
			Key:      "RequireMinCompatClient",
			OldValue: "luminous",
			Value:    "reef",
		},
	}, diff)
}

func (s *serviceTestSuite) TestDumpConfig() {
	s.cephMock.On("DumpConfig").Return(models.CephConfig{
		"osd": {
			"test_key": "value",
		},
	}, nil).Once()

	cfg, err := s.svc.DumpConfig(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(models.CephConfig{
		"osd": {
			"test_key": "value",
		},
	}, cfg)
}

func (s *serviceTestSuite) TestDumpOSDConfig() {
	s.cephMock.On("ClusterReport").Return(models.ClusterReport{
		AllowCrimson:           true,
		NearfullRatio:          0.85,
		BackfillfullRatio:      0.9,
		FullRatio:              0.95,
		RequireMinCompatClient: "reef",
	}, nil).Once()

	cfg, err := s.svc.DumpOSDConfig(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(models.CephOSDConfig{
		AllowCrimson:           true,
		NearfullRatio:          0.85,
		BackfillfullRatio:      0.9,
		FullRatio:              0.95,
		RequireMinCompatClient: "reef",
	}, cfg)
}

// Definitions ...

type serviceTestSuite struct {
	suite.Suite

	ctx        context.Context
	cancel     context.CancelFunc
	cephMock   *ceph.Mock
	differMock *differ.Mock
	svc        Service
}

func (s *serviceTestSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 3*time.Second)

	s.cephMock = ceph.NewMock()
	s.differMock = differ.NewMock()
	s.svc = New(s.cephMock, s.differMock)
}

func (s *serviceTestSuite) TearDownTest() {
	s.cephMock.AssertExpectations(s.T())
	s.differMock.AssertExpectations(s.T())

	s.svc = nil
	s.cephMock = nil

	s.cancel()

	s.ctx = nil
	s.cancel = nil
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, &serviceTestSuite{})
}

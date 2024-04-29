package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/teran/cephctl/ceph"
	"github.com/teran/cephctl/models"
	ptr "github.com/teran/go-ptr"
)

func (s *serviceTestSuite) TestApplyCephConfig() {
	s.cephMock.On("DumpConfig").Return(models.CephConfig{
		"osd": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "old_value",
		},
	}, nil).Once()

	s.cephMock.On("RemoveCephConfigOption", "osd", "test_key").Return(nil).Once()
	s.cephMock.On("ApplyCephConfigOption", "mon", "test_key", "value").Return(nil).Once()
	s.cephMock.On("ApplyCephConfigOption", "osd.3", "test_key", "value").Return(nil).Once()

	err := s.svc.ApplyCephConfig(s.ctx, models.CephConfig{
		"mon": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "value",
		},
	})
	s.Require().NoError(err)
}

func (s *serviceTestSuite) TestCheckClusterHealth() {
	s.cephMock.On("ClusterStatus").Return(models.ClusterStatus{
		HealthStatus: models.ClusterStatusHealthOK,
		MutedChecks:  []models.ClusterStatusMutedCheck{},
		QuorumAmount: 5,
	}, nil).Once()

	chi, err := s.svc.CheckClusterHealth(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal([]models.ClusterHealthIndicator{
		{
			Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
			CurrentValue:       "HEALTH_OK",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeQuorum,
			CurrentValue:       "5",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeMonsDown,
			CurrentValue:       "0",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeMgrsDown,
			CurrentValue:       "0",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
			CurrentValue:       "0",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeMDSsDown,
			CurrentValue:       "0",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
			CurrentValue:       "0",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeUncleanPGs,
			CurrentValue:       "0",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeInactivePGs,
			CurrentValue:       "0",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
	}, chi)
}

func (s *serviceTestSuite) TestDiffCephConfig() {
	s.cephMock.On("DumpConfig").Return(models.CephConfig{
		"osd": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "old_value",
		},
	}, nil).Once()

	diff, err := s.svc.DiffCephConfig(s.ctx, models.CephConfig{
		"mon": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "value",
		},
	})
	s.Require().NoError(err)
	s.Require().ElementsMatch([]models.CephConfigDifference{
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

// Definitions ...

type serviceTestSuite struct {
	suite.Suite

	ctx      context.Context
	cancel   context.CancelFunc
	cephMock *ceph.Mock
	svc      Service
}

func (s *serviceTestSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 3*time.Second)

	s.cephMock = ceph.NewMock()
	s.svc = New(s.cephMock)
}

func (s *serviceTestSuite) TearDownTest() {
	s.cephMock.AssertExpectations(s.T())

	s.svc = nil
	s.cephMock = nil

	s.cancel()

	s.ctx = nil
	s.cancel = nil
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, &serviceTestSuite{})
}

package ceph

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/teran/cephctl/models"
)

func init() {
	log.SetLevel(log.TraceLevel)
}

func TestApplyCephConfigOption(t *testing.T) {
	r := require.New(t)

	c := New("testdata/ceph_mock_ApplyCephConfigOption")
	err := c.ApplyCephConfigOption(context.Background(), "section", "key", "value")
	r.NoError(err)
}

func TestClusterReport(t *testing.T) {
	r := require.New(t)

	c := New("testdata/ceph_mock_ClusterReport")
	rep, err := c.ClusterReport(context.Background())
	r.NoError(err)
	r.Equal(models.ClusterReport{
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
		TotalOSDUsedDataKB: uint64(10_888_918_388),
		TotalOSDUsedMetaKB: uint64(450_830_044),
		TotalOSDUsedOMAPKB: uint64(5_881_251),
		NumPools:           14,
		NumPGs:             234,
		NumPGsByState: map[string]uint32{
			"active": 234,
			"clean":  234,
		},
	}, rep)
}

func TestClusterStatus(t *testing.T) {
	r := require.New(t)

	c := New("testdata/ceph_mock_ClusterStatus")
	st, err := c.ClusterStatus(context.Background())
	r.NoError(err)
	r.Equal(models.ClusterStatus{
		HealthStatus: models.ClusterStatusHealthWARN,
		Checks: []models.ClusterStatusCheck{
			{
				Code:     "OSD_NEARFULL",
				Severity: models.ClusterStatusHealthWARN,
				Summary:  "13 nearfull osd(s)",
			},
		},
		MutedChecks: []models.ClusterStatusMutedCheck{
			{
				Code:    "OSD_NEARFULL",
				Summary: "13 nearfull osd(s)",
			},
		},
		MonsTotal:      5,
		QuorumAmount:   5,
		MonsDownAmount: 0,
		MGRsDownAmount: 0,
		MDSsDownAmount: 0,
		OSDsDownAmount: 0,
	}, st)
}

func TestDumpConfig(t *testing.T) {
	r := require.New(t)

	c := New("testdata/ceph_mock_ConfigDumpParse")
	cfg, err := c.DumpConfig(context.Background())
	r.NoError(err)
	r.Equal(models.CephConfig{
		"client.radosgw": {
			"rgw_cache_lru_size": "100000",
		},
	}, cfg)
}

func TestListDevices(t *testing.T) {
	r := require.New(t)
	c := New("testdata/ceph_mock_ListDevices")
	devices, err := c.ListDevices(context.Background())
	r.NoError(err)
	r.Equal([]models.Device{
		{
			ID:        "CT4000P3SSD8_XXXYYYZZZ60B",
			Daemons:   []string{},
			WearLevel: 0.09000000357627869,
		},
		{
			ID:        "CT4000P3SSD8_XXXYYYZZZD0F",
			Daemons:   []string{"osd.7", "osd.8"},
			WearLevel: 0.10000000149011612,
		},
		{
			ID:        "CT4000P3SSD8_XXXYYYZZZFCE",
			Daemons:   []string{"osd.4", "osd.5"},
			WearLevel: 0.10999999940395355,
		},
	}, devices)
}

func TestRemoveCephConfigOption(t *testing.T) {
	r := require.New(t)

	c := New("testdata/ceph_mock_RemoveCephConfigOption")
	err := c.RemoveCephConfigOption(context.Background(), "section", "key")
	r.NoError(err)
}

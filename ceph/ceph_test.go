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
			{
				Code:     "POOL_NEARFULL",
				Severity: models.ClusterStatusHealthWARN,
				Summary:  "14 pool(s) nearfull",
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

func TestRemoveCephConfigOption(t *testing.T) {
	r := require.New(t)

	c := New("testdata/ceph_mock_RemoveCephConfigOption")
	err := c.RemoveCephConfigOption(context.Background(), "section", "key")
	r.NoError(err)
}

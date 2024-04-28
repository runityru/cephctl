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

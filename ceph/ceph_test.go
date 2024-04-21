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

func TestConfigDumpParse(t *testing.T) {
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

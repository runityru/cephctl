package cephconfig

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
)

func TestNew(t *testing.T) {
	r := require.New(t)

	cfg, err := New([]byte(`{"global": {"key":"value"}}`))
	r.NoError(err)
	r.Equal(models.CephConfig{
		"global": {
			"key": "value",
		},
	}, cfg)
}

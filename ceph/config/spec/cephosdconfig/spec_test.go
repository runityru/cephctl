package cephosdconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
)

func TestNewValidConfig(t *testing.T) {
	r := require.New(t)

	data, err := os.ReadFile("testdata/full.json")
	r.NoError(err)

	cfg, err := New(data)
	r.NoError(err)
	r.Equal(models.CephOSDConfig{
		AllowCrimson:           true,
		NearfullRatio:          0.75,
		BackfillfullRatio:      0.8,
		FullRatio:              0.85,
		RequireMinCompatClient: "squid",
	}, cfg)
}

func TestNewEmptyConfig(t *testing.T) {
	r := require.New(t)

	data, err := os.ReadFile("testdata/empty.json")
	r.NoError(err)

	cfg, err := New(data)
	r.NoError(err)
	r.Equal(models.CephOSDConfig{
		AllowCrimson:           false,
		NearfullRatio:          0.85,
		BackfillfullRatio:      0.9,
		FullRatio:              0.95,
		RequireMinCompatClient: "reef",
	}, cfg)
}

package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
)

func TestOSDsNumDaemonVersions(t *testing.T) {
	tcs := []testCase{
		{
			name: "single version",
			in: models.ClusterReport{
				NumOSDsByVersion: map[string]uint16{
					"18.2.2": 3,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsNumDaemonVersions,
				CurrentValue:       "1",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "two versions",
			in: models.ClusterReport{
				NumOSDsByVersion: map[string]uint16{
					"18.2.1": 1,
					"18.2.2": 3,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsNumDaemonVersions,
				CurrentValue:       "2",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
		{
			name: "three versions",
			in: models.ClusterReport{
				NumOSDsByVersion: map[string]uint16{
					"18.2.0": 1,
					"18.2.1": 2,
					"18.2.2": 3,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsNumDaemonVersions,
				CurrentValue:       "3",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
		{
			name: "four versions",
			in: models.ClusterReport{
				NumOSDsByVersion: map[string]uint16{
					"17.2.9": 4,
					"18.2.0": 1,
					"18.2.1": 2,
					"18.2.2": 3,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsNumDaemonVersions,
				CurrentValue:       "4",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
		{
			name: "no versions",
			in: models.ClusterReport{
				NumOSDsByVersion: map[string]uint16{},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsNumDaemonVersions,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusUnknown,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := OSDsNumDaemonVersions(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

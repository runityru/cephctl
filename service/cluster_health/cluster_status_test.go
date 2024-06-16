package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
)

func TestClusterStatus(t *testing.T) {
	tcs := []testCase{
		{
			name: "HEALTH_OK",
			in: models.ClusterReport{
				HealthStatus: models.ClusterStatusHealthOK,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
				CurrentValue:       "HEALTH_OK",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "HEALTH_WARN",
			in: models.ClusterReport{
				HealthStatus: models.ClusterStatusHealthWARN,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
				CurrentValue:       "HEALTH_WARN",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
		{
			name: "HEALTH_ERR",
			in: models.ClusterReport{
				HealthStatus: models.ClusterStatusHealthERR,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
				CurrentValue:       "HEALTH_ERR",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
		{
			name: "RANDOM_VALUE",
			in: models.ClusterReport{
				HealthStatus: models.ClusterStatusHealthUnknown,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
				CurrentValue:       "UNKNOWN",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusUnknown,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := ClusterStatus(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/runityru/cephctl/models"
)

func TestAllowCrimson(t *testing.T) {
	tcs := []testCase{
		{
			name: "crimson is allowed",
			in: models.ClusterReport{
				AllowCrimson: true,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeAllowCrimson,
				CurrentValue:       "true",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
		{
			name: "crimson is disallowed",
			in: models.ClusterReport{
				AllowCrimson: false,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeAllowCrimson,
				CurrentValue:       "false",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := AllowCrimson(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

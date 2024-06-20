package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
)

func TestOSDsOut(t *testing.T) {
	tcs := []testCase{
		{
			name: "all osds are in",
			in: models.ClusterReport{
				NumOSDs:   10,
				NumOSDsIn: 10,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsOut,
				CurrentValue:       "0 of 10",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some osds are out",
			in: models.ClusterReport{
				NumOSDs:   10,
				NumOSDsIn: 7,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsOut,
				CurrentValue:       "3 of 10",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := OSDsOut(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

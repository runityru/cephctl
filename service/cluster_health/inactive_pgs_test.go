package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
)

func TestInactivePGs(t *testing.T) {
	tcs := []testCase{
		{
			name: "no inactive pgs",
			in: models.ClusterReport{
				NumPGs: 10,
				NumPGsByState: map[string]uint32{
					"active": 10,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeInactivePGs,
				CurrentValue:       "0 of 10",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some inactive pgs",
			in: models.ClusterReport{
				NumPGs: 10,
				NumPGsByState: map[string]uint32{
					"active":   7,
					"clean":    7,
					"degraded": 3,
					"inactive": 3,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeInactivePGs,
				CurrentValue:       "3 of 10",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := InactivePGs(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

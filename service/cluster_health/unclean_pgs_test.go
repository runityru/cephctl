package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/runityru/cephctl/models"
)

func TestUncleanPGs(t *testing.T) {
	tcs := []testCase{
		{
			name: "no unclean pgs",
			in: models.ClusterReport{
				NumPGs: 10,
				NumPGsByState: map[string]uint32{
					"clean": 10,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeUncleanPGs,
				CurrentValue:       "0 of 10",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some unclean pgs",
			in: models.ClusterReport{
				NumPGs: 13,
				NumPGsByState: map[string]uint32{
					"clean":    10,
					"active":   13,
					"degraded": 3,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeUncleanPGs,
				CurrentValue:       "3 of 13",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := UncleanPGs(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

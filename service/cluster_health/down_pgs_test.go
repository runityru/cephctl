package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
)

func TestDownPGs(t *testing.T) {
	tcs := []testCase{
		{
			name: "no down OSDs",
			in: models.ClusterReport{
				NumPGs: 10,
				NumPGsByState: map[string]uint32{
					"active": 10,
					"clean":  10,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeDownPGs,
				CurrentValue:       "0 of 10",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some down OSDs",
			in: models.ClusterReport{
				NumPGs: 10,
				NumPGsByState: map[string]uint32{
					"active": 3,
					"down":   5,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeDownPGs,
				CurrentValue:       "5 of 10",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := DownPGs(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

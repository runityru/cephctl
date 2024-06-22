package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
)

func TestMutesAmount(t *testing.T) {
	tcs := []testCase{
		{
			name: "no mutes",
			in: models.ClusterReport{
				MutedChecks: []models.ClusterStatusMutedCheck{},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
				CurrentValue:       "0 of 0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "mute present",
			in: models.ClusterReport{
				MutedChecks: []models.ClusterStatusMutedCheck{
					{
						Code:    "SOME_CHECK",
						Summary: "There a check failed, beware!",
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
				CurrentValue:       "1 of 0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := MutesAmount(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

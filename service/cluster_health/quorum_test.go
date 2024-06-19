package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/runityru/cephctl/models"
)

func TestQuorum(t *testing.T) {
	tcs := []testCase{
		{
			name: "all in quorum",
			in: models.ClusterReport{
				NumMons:         5,
				NumMonsInQuorum: 5,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeQuorum,
				CurrentValue:       "5 of 5",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some out of quorum",
			in: models.ClusterReport{
				NumMons:         5,
				NumMonsInQuorum: 3,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeQuorum,
				CurrentValue:       "3 of 5",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := Quorum(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

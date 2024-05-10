package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
)

type testCase struct {
	name   string
	in     models.ClusterReport
	expOut models.ClusterHealthIndicator
}

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
				CurrentValue:       "5",
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
				CurrentValue:       "3",
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

func TestOSDsDown(t *testing.T) {
	tcs := []testCase{
		{
			name: "all osds are alive",
			in: models.ClusterReport{
				NumOSDs:   10,
				NumOSDsUp: 10,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some osds are down",
			in: models.ClusterReport{
				NumOSDs:   10,
				NumOSDsUp: 7,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
				CurrentValue:       "3",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := OSDsDown(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

func TestMutesAmount(t *testing.T) {
	tcs := []testCase{
		{
			name: "no mutes",
			in: models.ClusterReport{
				MutedChecks: []models.ClusterStatusMutedCheck{},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
				CurrentValue:       "0",
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
				CurrentValue:       "1",
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
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some unclean pgs",
			in: models.ClusterReport{
				NumPGs: 10,
				NumPGsByState: map[string]uint32{
					"active":   10,
					"clean":    7,
					"degraded": 3,
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeUncleanPGs,
				CurrentValue:       "3",
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
				CurrentValue:       "0",
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
				CurrentValue:       "3",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
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

package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
)

type testCase struct {
	name   string
	in     models.ClusterStatus
	expOut models.ClusterHealthIndicator
}

func TestClusterStatus(t *testing.T) {
	tcs := []testCase{
		{
			name: "HEALTH_OK",
			in: models.ClusterStatus{
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
			in: models.ClusterStatus{
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
			in: models.ClusterStatus{
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
			in: models.ClusterStatus{
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
			in: models.ClusterStatus{
				QuorumAmount: 5,
				MonsTotal:    5,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeQuorum,
				CurrentValue:       "5",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some out of quorum",
			in: models.ClusterStatus{
				QuorumAmount: 3,
				MonsTotal:    5,
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

func TestMonsDown(t *testing.T) {
	tcs := []testCase{
		{
			name: "all mons are alive",
			in: models.ClusterStatus{
				MonsDownAmount: 0,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMonsDown,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some monitors are down",
			in: models.ClusterStatus{
				MonsDownAmount: 1,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMonsDown,
				CurrentValue:       "1",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := MonsDown(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

func TestMgrsDown(t *testing.T) {
	tcs := []testCase{
		{
			name: "all mgrs are alive",
			in: models.ClusterStatus{
				MGRsDownAmount: 0,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMgrsDown,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some mgrs are down",
			in: models.ClusterStatus{
				MGRsDownAmount: 3,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMgrsDown,
				CurrentValue:       "3",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := MgrsDown(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

func TestOSDsDown(t *testing.T) {
	tcs := []testCase{
		{
			name: "all osds are alive",
			in: models.ClusterStatus{
				OSDsDownAmount: 0,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some osds are down",
			in: models.ClusterStatus{
				OSDsDownAmount: 3,
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

func TestMDSsDown(t *testing.T) {
	tcs := []testCase{
		{
			name: "all mds are alive",
			in: models.ClusterStatus{
				MDSsDownAmount: 0,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMDSsDown,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some mds are down",
			in: models.ClusterStatus{
				MDSsDownAmount: 3,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeMDSsDown,
				CurrentValue:       "3",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := MDSsDown(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

func TestMutesAmount(t *testing.T) {
	tcs := []testCase{
		{
			name: "no mutes",
			in: models.ClusterStatus{
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
			in: models.ClusterStatus{
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
			in: models.ClusterStatus{
				UncleanPGs: 0,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeUncleanPGs,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some unclean pgs",
			in: models.ClusterStatus{
				UncleanPGs: 3,
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
			in: models.ClusterStatus{
				InactivePGs: 0,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeInactivePGs,
				CurrentValue:       "0",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "some inactive pgs",
			in: models.ClusterStatus{
				InactivePGs: 3,
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

package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
)

func TestIPCollision(t *testing.T) {
	tcs := []testCase{
		{
			name: "no collisions",
			in: models.ClusterReport{
				OSDDaemons: []models.OSDDaemon{
					{
						Hostname: "host1",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.1",
					},
					{
						Hostname: "host1",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.1",
					},
					{
						Hostname: "host2",
						FrontIP:  "192.168.1.2",
						BackIP:   "192.168.2.2",
					},
					{
						Hostname: "host2",
						FrontIP:  "192.168.1.2",
						BackIP:   "192.168.2.2",
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeIPCollision,
				CurrentValue:       "all hosts have their own IPs",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "front IP collision",
			in: models.ClusterReport{
				OSDDaemons: []models.OSDDaemon{
					{
						Hostname: "host1",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.1",
					},
					{
						Hostname: "host1",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.1",
					},
					{
						Hostname: "host2",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.2",
					},
					{
						Hostname: "host2",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.2",
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeIPCollision,
				CurrentValue:       "2 or more hosts have the same front IP",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
		{
			name: "back IP collision",
			in: models.ClusterReport{
				OSDDaemons: []models.OSDDaemon{
					{
						Hostname: "host1",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.1",
					},
					{
						Hostname: "host1",
						FrontIP:  "192.168.1.1",
						BackIP:   "192.168.2.1",
					},
					{
						Hostname: "host2",
						FrontIP:  "192.168.1.2",
						BackIP:   "192.168.2.1",
					},
					{
						Hostname: "host2",
						FrontIP:  "192.168.1.2",
						BackIP:   "192.168.2.1",
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeIPCollision,
				CurrentValue:       "2 or more hosts have the same back IP",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := IPCollision(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

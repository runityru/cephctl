package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/runityru/cephctl/models"
)

func TestDeviceHealth(t *testing.T) {
	tcs := []testCase{
		{
			name: "All OK",
			in: models.ClusterReport{
				Devices: []models.Device{
					{
						ID:        "1",
						Daemons:   []string{"osd.1"},
						WearLevel: 0.1,
					},
					{
						ID:        "2",
						Daemons:   []string{"osd.2"},
						WearLevel: 0.05,
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeDeviceHealthWearout,
				CurrentValue:       ">50.0%: 0 device(s); >75.0%: 0 device(s)",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "Proper status on multiple conditions (dangerous)",
			in: models.ClusterReport{
				Devices: []models.Device{
					{
						ID:        "0",
						Daemons:   []string{"osd.0"},
						WearLevel: 0.1,
					},
					{
						ID:        "1",
						Daemons:   []string{"osd.1"},
						WearLevel: 0.51,
					},
					{
						ID:        "2",
						Daemons:   []string{"osd.2"},
						WearLevel: 0.751,
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeDeviceHealthWearout,
				CurrentValue:       ">50.0%: 1 device(s); >75.0%: 1 device(s)",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
		{
			name: "Proper status on multiple conditions (at risk)",
			in: models.ClusterReport{
				Devices: []models.Device{
					{
						ID:        "0",
						Daemons:   []string{"osd.0"},
						WearLevel: 0.1,
					},
					{
						ID:        "1",
						Daemons:   []string{"osd.1"},
						WearLevel: 0.51,
					},
					{
						ID:        "2",
						Daemons:   []string{"osd.2"},
						WearLevel: 0.55,
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeDeviceHealthWearout,
				CurrentValue:       ">50.0%: 2 device(s); >75.0%: 0 device(s)",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
		{
			name: "Devices without daemons",
			in: models.ClusterReport{
				Devices: []models.Device{
					{
						ID:        "0",
						Daemons:   []string{"osd.0"},
						WearLevel: 0.1,
					},
					{
						ID:        "1",
						Daemons:   []string{"osd.1"},
						WearLevel: 0.51,
					},
					{
						ID:        "2",
						Daemons:   []string{},
						WearLevel: 0.99,
					},
				},
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeDeviceHealthWearout,
				CurrentValue:       ">50.0%: 1 device(s); >75.0%: 0 device(s)",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := DeviceHealth(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

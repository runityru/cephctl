package cluster_health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
)

func TestOSDsMetadataSize(t *testing.T) {
	tcs := []testCase{
		{
			name: "metadata size is <7%",
			in: models.ClusterReport{
				TotalOSDCapacityKB: 10000,
				TotalOSDUsedMetaKB: 100,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsMetadataSize,
				CurrentValue:       "1.00%",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
			},
		},
		{
			name: "metadata size is >7%",
			in: models.ClusterReport{
				TotalOSDCapacityKB: 10000,
				TotalOSDUsedMetaKB: 1782,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsMetadataSize,
				CurrentValue:       "17.82%",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
			},
		},
		{
			name: "metadata size is >10%",
			in: models.ClusterReport{
				TotalOSDCapacityKB: 10000,
				TotalOSDUsedMetaKB: 2006,
			},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsMetadataSize,
				CurrentValue:       "20.06%",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			},
		},
		{
			name: "empty structure (division by zero provocation)",
			in:   models.ClusterReport{},
			expOut: models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeOSDsMetadataSize,
				CurrentValue:       "NaN%",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusUnknown,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			i, err := OSDsMetadataSize(context.Background(), tc.in)
			r.NoError(err)
			r.Equal(tc.expOut, i)
		})
	}
}

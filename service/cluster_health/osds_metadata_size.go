package cluster_health

import (
	"context"
	"strconv"

	"github.com/teran/cephctl/models"
)

func OSDsMetadataSize(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusUnknown

	metadataSizePercentage := 100.0 / float64(cr.TotalOSDCapacityKB) * float64(cr.TotalOSDUsedMetaKB)

	if metadataSizePercentage > 20.0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	} else if metadataSizePercentage > 15.0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	} else if metadataSizePercentage > 0 {
		st = models.ClusterHealthIndicatorStatusGood
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsMetadataSize,
		CurrentValue:       strconv.FormatFloat(metadataSizePercentage, 'f', 2, 64) + "%",
		CurrentValueStatus: st,
	}, nil
}

package cluster_health

import (
	"context"
	"fmt"

	"github.com/teran/cephctl/models"
)

func MutesAmount(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	if len(cr.MutedChecks) > 0 {
		return models.ClusterHealthIndicator{
			Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
			CurrentValue:       fmt.Sprintf("%d of %d", len(cr.MutedChecks), len(cr.Checks)),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
		}, nil
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
		CurrentValue:       "0 of 0",
		CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
	}, nil
}

func OSDsDown(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	numOSDsDown := cr.NumOSDs - cr.NumOSDsUp
	if numOSDsDown > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
		CurrentValue:       fmt.Sprintf("%d of %d", numOSDsDown, cr.NumOSDs),
		CurrentValueStatus: st,
	}, nil
}

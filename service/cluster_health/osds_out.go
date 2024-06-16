package cluster_health

import (
	"context"
	"fmt"

	"github.com/teran/cephctl/models"
)

func OSDsOut(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	numOSDsOut := cr.NumOSDs - cr.NumOSDsIn
	if numOSDsOut > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsOut,
		CurrentValue:       fmt.Sprintf("%d of %d", numOSDsOut, cr.NumOSDs),
		CurrentValueStatus: st,
	}, nil
}

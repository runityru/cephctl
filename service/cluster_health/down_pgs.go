package cluster_health

import (
	"context"
	"fmt"

	"github.com/runityru/cephctl/models"
)

func DownPGs(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood

	downPGs := cr.NumPGsByState["down"]
	if downPGs > 0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeDownPGs,
		CurrentValue:       fmt.Sprintf("%d of %d", downPGs, cr.NumPGs),
		CurrentValueStatus: st,
	}, nil
}

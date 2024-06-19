package cluster_health

import (
	"context"
	"fmt"

	"github.com/runityru/cephctl/models"
)

func InactivePGs(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood

	activePGs, _ := cr.NumPGsByState["active"]
	inactivePGs := cr.NumPGs - activePGs
	if inactivePGs > 0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeInactivePGs,
		CurrentValue:       fmt.Sprintf("%d of %d", inactivePGs, cr.NumPGs),
		CurrentValueStatus: st,
	}, nil
}

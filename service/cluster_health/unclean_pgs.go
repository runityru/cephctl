package cluster_health

import (
	"context"
	"fmt"

	"github.com/runityru/cephctl/models"
)

func UncleanPGs(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood

	cleanPGs := cr.NumPGsByState["clean"]
	uncleanPGs := cr.NumPGs - cleanPGs
	if uncleanPGs > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeUncleanPGs,
		CurrentValue:       fmt.Sprintf("%d of %d", uncleanPGs, cr.NumPGs),
		CurrentValueStatus: st,
	}, nil
}

package cluster_health

import (
	"context"
	"fmt"

	"github.com/runityru/cephctl/models"
)

func Quorum(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cr.NumMonsInQuorum < cr.NumMons {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeQuorum,
		CurrentValue:       fmt.Sprintf("%d of %d", cr.NumMonsInQuorum, cr.NumMons),
		CurrentValueStatus: st,
	}, nil
}

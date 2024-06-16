package cluster_health

import (
	"context"
	"strconv"

	"github.com/teran/cephctl/models"
)

func AllowCrimson(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cr.AllowCrimson {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeAllowCrimson,
		CurrentValue:       strconv.FormatBool(cr.AllowCrimson),
		CurrentValueStatus: st,
	}, nil
}

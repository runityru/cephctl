package cluster_health

import (
	"context"

	"github.com/teran/cephctl/models"
)

func ClusterStatus(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	const indicator = models.ClusterHealthIndicatorTypeClusterStatus

	switch cr.HealthStatus {
	case models.ClusterStatusHealthOK:
		return models.ClusterHealthIndicator{
			Indicator:          indicator,
			CurrentValue:       string(cr.HealthStatus),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		}, nil
	case models.ClusterStatusHealthWARN:
		return models.ClusterHealthIndicator{
			Indicator:          indicator,
			CurrentValue:       string(cr.HealthStatus),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
		}, nil
	case models.ClusterStatusHealthERR:
		return models.ClusterHealthIndicator{
			Indicator:          indicator,
			CurrentValue:       string(cr.HealthStatus),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
		}, nil
	}

	return models.ClusterHealthIndicator{
		Indicator:          indicator,
		CurrentValue:       string(cr.HealthStatus),
		CurrentValueStatus: models.ClusterHealthIndicatorStatusUnknown,
	}, nil
}

package cluster_health

import (
	"context"
	"strconv"

	"github.com/teran/cephctl/models"
)

func OSDsNumDaemonVersions(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	numVersions := len(cr.NumOSDsByVersion)

	st := models.ClusterHealthIndicatorStatusUnknown
	if numVersions > 2 {
		st = models.ClusterHealthIndicatorStatusDangerous
	} else if numVersions == 2 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	} else if numVersions == 1 {
		st = models.ClusterHealthIndicatorStatusGood
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsNumDaemonVersions,
		CurrentValue:       strconv.FormatInt(int64(numVersions), 10),
		CurrentValueStatus: st,
	}, nil
}

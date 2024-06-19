package cluster_health

import (
	"context"

	"github.com/runityru/cephctl/models"
)

func IPCollision(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	frontIPMap := make(map[string]map[string]struct{})
	backIPMap := make(map[string]map[string]struct{})

	for _, osd := range cr.OSDDaemons {
		if _, ok := frontIPMap[osd.FrontIP]; !ok {
			frontIPMap[osd.FrontIP] = map[string]struct{}{
				osd.Hostname: {},
			}
		} else {
			frontIPMap[osd.FrontIP][osd.Hostname] = struct{}{}
		}

		if _, ok := backIPMap[osd.BackIP]; !ok {
			backIPMap[osd.BackIP] = map[string]struct{}{
				osd.Hostname: {},
			}
		} else {
			backIPMap[osd.BackIP][osd.Hostname] = struct{}{}
		}
	}

	for _, v := range frontIPMap {
		if len(v) > 1 {
			return models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeIPCollision,
				CurrentValue:       "2 or more hosts have the same front IP",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			}, nil
		}
	}

	for _, v := range backIPMap {
		if len(v) > 1 {
			return models.ClusterHealthIndicator{
				Indicator:          models.ClusterHealthIndicatorTypeIPCollision,
				CurrentValue:       "2 or more hosts have the same back IP",
				CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
			}, nil
		}
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeIPCollision,
		CurrentValue:       "all hosts have their own IPs",
		CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
	}, nil
}

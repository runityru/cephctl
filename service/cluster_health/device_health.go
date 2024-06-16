package cluster_health

import (
	"context"
	"fmt"

	"github.com/teran/cephctl/models"
)

func DeviceHealth(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	const (
		riskLevel      = 0.5
		dangerousLevel = 0.75
	)

	st := models.ClusterHealthIndicatorStatusGood

	var (
		atRiskDevs      uint16
		atDangerousDevs uint16
	)

	for _, dev := range cr.Devices {
		if len(dev.Daemons) > 0 {
			if dev.WearLevel > dangerousLevel {
				atDangerousDevs++
			} else if dev.WearLevel > riskLevel {
				atRiskDevs++
			}
		}
	}

	if atDangerousDevs > 0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	} else if atRiskDevs > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeDeviceHealthWearout,
		CurrentValue:       fmt.Sprintf(">%.1f%%: %d device(s); >%.1f%%: %d device(s)", riskLevel*100, atRiskDevs, dangerousLevel*100, atDangerousDevs),
		CurrentValueStatus: st,
	}, nil
}

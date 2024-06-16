package healthcheck

import (
	"context"

	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/printer"
	"github.com/teran/cephctl/service"
	clusterHealth "github.com/teran/cephctl/service/cluster_health"
)

type HealthcheckConfig struct {
	Service service.Service
	Printer printer.Printer
}

func Healthcheck(ctx context.Context, hc HealthcheckConfig) error {
	indicators, err := hc.Service.CheckClusterHealth(ctx, []clusterHealth.ClusterHealthCheck{
		clusterHealth.ClusterStatus,
		clusterHealth.Quorum,
		clusterHealth.OSDsDown,
		clusterHealth.OSDsOut,
		clusterHealth.MutesAmount,
		clusterHealth.DownPGs,
		clusterHealth.UncleanPGs,
		clusterHealth.InactivePGs,
		clusterHealth.AllowCrimson,
		clusterHealth.OSDsMetadataSize,
		clusterHealth.OSDsNumDaemonVersions,
		clusterHealth.IPCollision,
		clusterHealth.DeviceHealth,
	})
	if err != nil {
		return err
	}

	for _, indicator := range indicators {
		var printFn func(string, ...any) = hc.Printer.HiRed

		switch indicator.CurrentValueStatus {
		case models.ClusterHealthIndicatorStatusGood:
			printFn = hc.Printer.Green
		case models.ClusterHealthIndicatorStatusAtRisk:
			printFn = hc.Printer.Yellow
		case models.ClusterHealthIndicatorStatusDangerous:
			printFn = hc.Printer.Red
		}

		printFn(
			"[%s] %s = %s",
			padTo(
				string(indicator.CurrentValueStatus),
				len(string(models.ClusterHealthIndicatorStatusDangerous)),
			), indicator.Indicator, indicator.CurrentValue,
		)
	}
	return nil
}

func padTo(s string, n int) string {
	for len(s) < n {
		s = " " + s
	}
	return s
}

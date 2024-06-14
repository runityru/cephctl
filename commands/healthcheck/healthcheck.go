package healthcheck

import (
	"context"

	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/printer"
	"github.com/teran/cephctl/service"
)

type HealthcheckConfig struct {
	Service service.Service
	Printer printer.Printer
}

func Healthcheck(ctx context.Context, hc HealthcheckConfig) error {
	indicators, err := hc.Service.CheckClusterHealth(ctx)
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

package status

import (
	"context"

	"github.com/fatih/color"
	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/service"
)

func Healthcheck(ctx context.Context, svc service.Service) error {
	indicators, err := svc.CheckClusterHealth(ctx)
	if err != nil {
		return err
	}

	for _, indicator := range indicators {
		var printFn func(string, ...interface{}) = color.HiRed
		switch indicator.CurrentValueStatus {
		case models.ClusterHealthIndicatorStatusGood:
			printFn = color.Green
		case models.ClusterHealthIndicatorStatusAtRisk:
			printFn = color.Yellow
		case models.ClusterHealthIndicatorStatusDangerous:
			printFn = color.Red
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

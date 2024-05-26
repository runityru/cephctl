package cluster_health

import (
	"context"
	"fmt"
	"strconv"

	"github.com/teran/cephctl/models"
)

type ClusterHealthCheck func(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error)

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

func DownPGs(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood

	downPGs, _ := cr.NumPGsByState["down"]
	if downPGs > 0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeDownPGs,
		CurrentValue:       fmt.Sprintf("%d of %d", downPGs, cr.NumPGs),
		CurrentValueStatus: st,
	}, nil
}

func InactivePGs(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood

	activePGs, _ := cr.NumPGsByState["active"]
	inactivePGs := cr.NumPGs - activePGs
	if inactivePGs > 0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeInactivePGs,
		CurrentValue:       fmt.Sprintf("%d of %d", inactivePGs, cr.NumPGs),
		CurrentValueStatus: st,
	}, nil
}

func MutesAmount(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	if len(cr.MutedChecks) > 0 {
		return models.ClusterHealthIndicator{
			Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
			CurrentValue:       fmt.Sprintf("%d of %d", len(cr.MutedChecks), len(cr.Checks)),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
		}, nil
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
		CurrentValue:       "0 of 0",
		CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
	}, nil
}

func OSDsDown(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	numOSDsDown := cr.NumOSDs - cr.NumOSDsUp
	if numOSDsDown > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
		CurrentValue:       fmt.Sprintf("%d of %d", numOSDsDown, cr.NumOSDs),
		CurrentValueStatus: st,
	}, nil
}

func OSDsMetadataSize(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusUnknown

	metadataSizePercentage := 100.0 / float64(cr.TotalOSDCapacityKB) * float64(cr.TotalOSDUsedMetaKB)

	if metadataSizePercentage > 20.0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	} else if metadataSizePercentage > 15.0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	} else if metadataSizePercentage > 0 {
		st = models.ClusterHealthIndicatorStatusGood
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsMetadataSize,
		CurrentValue:       strconv.FormatFloat(metadataSizePercentage, 'f', 2, 64) + "%",
		CurrentValueStatus: st,
	}, nil
}

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

func OSDsOut(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	numOSDsOut := cr.NumOSDs - cr.NumOSDsIn
	if numOSDsOut > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsOut,
		CurrentValue:       fmt.Sprintf("%d of %d", numOSDsOut, cr.NumOSDs),
		CurrentValueStatus: st,
	}, nil
}

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

func UncleanPGs(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood

	cleanPGs, _ := cr.NumPGsByState["clean"]
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

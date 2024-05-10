package cluster_health

import (
	"context"
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

func InactivePGs(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood

	activePGs, _ := cr.NumPGsByState["active"]
	inactivePGs := cr.NumPGs - activePGs
	if inactivePGs > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeInactivePGs,
		CurrentValue:       strconv.FormatUint(uint64(inactivePGs), 10),
		CurrentValueStatus: st,
	}, nil
}

func MutesAmount(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	if len(cr.MutedChecks) > 0 {
		return models.ClusterHealthIndicator{
			Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
			CurrentValue:       strconv.FormatInt(int64(len(cr.MutedChecks)), 10),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
		}, nil
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
		CurrentValue:       "0",
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
		CurrentValue:       strconv.FormatUint(uint64(numOSDsDown), 10),
		CurrentValueStatus: st,
	}, nil
}

func OSDsMetadataSize(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusUnknown

	metadataSizePercentage := 100.0 / float64(cr.TotalOSDCapacityKB) * float64(cr.TotalOSDUsedMetaKB)

	if metadataSizePercentage > 10.0 {
		st = models.ClusterHealthIndicatorStatusDangerous
	} else if metadataSizePercentage > 7.0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	} else if metadataSizePercentage > 0 {
		st = models.ClusterHealthIndicatorStatusGood
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsMetadataSize,
		CurrentValue:       strconv.FormatFloat(metadataSizePercentage, 'f', 2, 64),
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

func Quorum(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cr.NumMonsInQuorum < cr.NumMons {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeQuorum,
		CurrentValue:       strconv.FormatUint(uint64(cr.NumMonsInQuorum), 10),
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
		CurrentValue:       strconv.FormatUint(uint64(uncleanPGs), 10),
		CurrentValueStatus: st,
	}, nil
}

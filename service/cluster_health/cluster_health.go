package cluster_health

import (
	"context"
	"strconv"

	"github.com/teran/cephctl/models"
)

type ClusterHealthCheck func(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error)

func ClusterStatus(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	const indicator = models.ClusterHealthIndicatorTypeClusterStatus

	switch cs.HealthStatus {
	case models.ClusterStatusHealthOK:
		return models.ClusterHealthIndicator{
			Indicator:          indicator,
			CurrentValue:       string(cs.HealthStatus),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		}, nil
	case models.ClusterStatusHealthWARN:
		return models.ClusterHealthIndicator{
			Indicator:          indicator,
			CurrentValue:       string(cs.HealthStatus),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
		}, nil
	case models.ClusterStatusHealthERR:
		return models.ClusterHealthIndicator{
			Indicator:          indicator,
			CurrentValue:       string(cs.HealthStatus),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusDangerous,
		}, nil
	}

	return models.ClusterHealthIndicator{
		Indicator:          indicator,
		CurrentValue:       string(cs.HealthStatus),
		CurrentValueStatus: models.ClusterHealthIndicatorStatusUnknown,
	}, nil
}

func Quorum(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cs.QuorumAmount < cs.MonsTotal {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeQuorum,
		CurrentValue:       strconv.FormatUint(uint64(cs.QuorumAmount), 10),
		CurrentValueStatus: st,
	}, nil
}

func MonsDown(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cs.MonsDownAmount > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeMonsDown,
		CurrentValue:       strconv.FormatUint(uint64(cs.MonsDownAmount), 10),
		CurrentValueStatus: st,
	}, nil
}

func MgrsDown(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cs.MGRsDownAmount > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeMgrsDown,
		CurrentValue:       strconv.FormatUint(uint64(cs.MGRsDownAmount), 10),
		CurrentValueStatus: st,
	}, nil
}

func OSDsDown(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cs.OSDsDownAmount > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
		CurrentValue:       strconv.FormatUint(uint64(cs.OSDsDownAmount), 10),
		CurrentValueStatus: st,
	}, nil
}

func RGWsDown(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cs.RGWsDownAmount > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeRGWsDown,
		CurrentValue:       strconv.FormatUint(uint64(cs.RGWsDownAmount), 10),
		CurrentValueStatus: st,
	}, nil
}

func MDSsDown(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	st := models.ClusterHealthIndicatorStatusGood
	if cs.MDSsDownAmount > 0 {
		st = models.ClusterHealthIndicatorStatusAtRisk
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeMDSsDown,
		CurrentValue:       strconv.FormatUint(uint64(cs.MDSsDownAmount), 10),
		CurrentValueStatus: st,
	}, nil
}

func MutesAmount(ctx context.Context, cs models.ClusterStatus) (models.ClusterHealthIndicator, error) {
	if len(cs.MutedChecks) > 0 {
		return models.ClusterHealthIndicator{
			Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
			CurrentValue:       strconv.FormatInt(int64(len(cs.MutedChecks)), 10),
			CurrentValueStatus: models.ClusterHealthIndicatorStatusAtRisk,
		}, nil
	}

	return models.ClusterHealthIndicator{
		Indicator:          models.ClusterHealthIndicatorTypeMutesAmount,
		CurrentValue:       "0",
		CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
	}, nil
}

package service

import (
	"context"

	"github.com/pkg/errors"
	diff "github.com/r3labs/diff/v3"
	log "github.com/sirupsen/logrus"
	ptr "github.com/teran/go-ptr"

	"github.com/teran/cephctl/ceph"
	"github.com/teran/cephctl/models"
	clusterHealth "github.com/teran/cephctl/service/cluster_health"
)

type Service interface {
	ApplyCephConfig(ctx context.Context, cfg models.CephConfig) error
	DiffCephConfig(ctx context.Context, cfg models.CephConfig) ([]models.CephConfigDifference, error)
	CheckClusterHealth(ctx context.Context) ([]models.ClusterHealthIndicator, error)
	DumpConfig(ctx context.Context) (models.CephConfig, error)
}

type service struct {
	c ceph.Ceph
}

func New(c ceph.Ceph) Service {
	return &service{
		c: c,
	}
}

func (s *service) ApplyCephConfig(ctx context.Context, cfg models.CephConfig) error {
	changes, err := s.DiffCephConfig(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "error comparing current and desired configuration")
	}

	log.Tracef("changelog: %#v", changes)

	for _, change := range changes {
		switch change.Kind {
		case models.CephConfigDifferenceKindRemove:
			if err := s.c.RemoveCephConfigOption(ctx, change.Section, change.Key); err != nil {
				return err
			}
		case models.CephConfigDifferenceKindAdd, models.CephConfigDifferenceKindChange:
			if err := s.c.ApplyCephConfigOption(ctx, change.Section, change.Key, *change.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *service) CheckClusterHealth(ctx context.Context) ([]models.ClusterHealthIndicator, error) {
	cs, err := s.c.ClusterStatus(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving cluster status")
	}

	checks := []clusterHealth.ClusterHealthCheck{
		clusterHealth.ClusterStatus,
		clusterHealth.Quorum,
		clusterHealth.MonsDown,
		clusterHealth.MgrsDown,
		clusterHealth.OSDsDown,
		clusterHealth.RGWsDown,
		clusterHealth.MDSsDown,
		clusterHealth.MutesAmount,
	}

	indicators := []models.ClusterHealthIndicator{}
	for _, checkFunc := range checks {
		indicator, err := checkFunc(ctx, cs)
		if err != nil {
			return nil, err
		}

		indicators = append(indicators, indicator)
	}

	return indicators, nil
}

func (s *service) DiffCephConfig(ctx context.Context, cfg models.CephConfig) ([]models.CephConfigDifference, error) {
	src, err := s.c.DumpConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving current configuration")
	}

	changelog, err := diff.Diff(src, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error comparing current and desired configuration")
	}

	changes := []models.CephConfigDifference{}
	for _, change := range changelog {
		switch change.Type {
		case diff.CREATE:
			v, ok := change.To.(string)
			if !ok {
				log.Warnf("unexpected value type: expected string, got %T", v)
				break
			}

			changes = append(changes, models.CephConfigDifference{
				Kind:    models.CephConfigDifferenceKindAdd,
				Section: change.Path[0],
				Key:     change.Path[1],
				Value:   ptr.String(v),
			})

		case diff.UPDATE:
			oldV, ok := change.From.(string)
			if !ok {
				log.Warnf("unexpected old value type: expected string, got %T", oldV)
				break
			}

			v, ok := change.To.(string)
			if !ok {
				log.Warnf("unexpected new value type: expected string, got %T", v)
				break
			}

			changes = append(changes, models.CephConfigDifference{
				Kind:     models.CephConfigDifferenceKindChange,
				Section:  change.Path[0],
				Key:      change.Path[1],
				OldValue: ptr.String(oldV),
				Value:    ptr.String(v),
			})

		case diff.DELETE:
			changes = append(changes, models.CephConfigDifference{
				Kind:    models.CephConfigDifferenceKindRemove,
				Section: change.Path[0],
				Key:     change.Path[1],
			})

		}
	}

	return changes, nil
}

func (s *service) DumpConfig(ctx context.Context) (models.CephConfig, error) {
	return s.c.DumpConfig(ctx)
}

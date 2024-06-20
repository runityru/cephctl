package service

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/runityru/cephctl/models"
	clusterHealth "github.com/runityru/cephctl/service/cluster_health"
)

var _ Service = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) ApplyCephConfig(_ context.Context, cfg models.CephConfig) error {
	args := m.Called(cfg)
	return args.Error(0)
}

func (m *Mock) DiffCephConfig(_ context.Context, cfg models.CephConfig) ([]models.CephConfigDifference, error) {
	args := m.Called(cfg)
	return args.Get(0).([]models.CephConfigDifference), args.Error(1)
}

func (m *Mock) CheckClusterHealth(context.Context, []clusterHealth.ClusterHealthCheck) ([]models.ClusterHealthIndicator, error) {
	args := m.Called()
	return args.Get(0).([]models.ClusterHealthIndicator), args.Error(1)
}

func (m *Mock) DumpConfig(context.Context) (models.CephConfig, error) {
	args := m.Called()
	return args.Get(0).(models.CephConfig), args.Error(1)
}

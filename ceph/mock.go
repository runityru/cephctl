package ceph

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/runityru/cephctl/models"
)

var _ Ceph = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) ApplyCephConfigOption(ctx context.Context, section, key, value string) error {
	args := m.Called(section, key, value)
	return args.Error(0)
}

func (m *Mock) ApplyCephOSDConfigOption(_ context.Context, key, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *Mock) ClusterReport(ctx context.Context) (models.ClusterReport, error) {
	args := m.Called()
	return args.Get(0).(models.ClusterReport), args.Error(1)
}

func (m *Mock) ClusterStatus(ctx context.Context) (models.ClusterStatus, error) {
	args := m.Called()
	return args.Get(0).(models.ClusterStatus), args.Error(1)
}

func (m *Mock) DumpConfig(_ context.Context) (models.CephConfig, error) {
	args := m.Called()
	return args.Get(0).(models.CephConfig), args.Error(1)
}

func (m *Mock) ListDevices(_ context.Context) ([]models.Device, error) {
	args := m.Called()
	return args.Get(0).([]models.Device), args.Error(1)
}

func (m *Mock) RemoveCephConfigOption(ctx context.Context, section, key string) error {
	args := m.Called(section, key)
	return args.Error(0)
}

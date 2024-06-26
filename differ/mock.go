package differ

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/runityru/cephctl/models"
)

var _ Differ = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) DiffCephConfig(ctx context.Context, from, to models.CephConfig) ([]models.CephConfigDifference, error) {
	args := m.Called(from, to)
	return args.Get(0).([]models.CephConfigDifference), args.Error(1)
}

func (m *Mock) DiffCephOSDConfig(ctx context.Context, from, to models.CephOSDConfig) ([]models.CephOSDConfigDifference, error) {
	args := m.Called(from, to)
	return args.Get(0).([]models.CephOSDConfigDifference), args.Error(1)
}

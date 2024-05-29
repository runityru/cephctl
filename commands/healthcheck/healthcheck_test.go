package healthcheck

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/printer"
	"github.com/teran/cephctl/service"
)

func TestHealthcheck(t *testing.T) {
	r := require.New(t)

	m := service.NewMock()
	defer m.AssertExpectations(t)

	p := printer.NewMock()
	defer p.AssertExpectations(t)

	m.On("CheckClusterHealth").Return([]models.ClusterHealthIndicator{
		{
			Indicator:          models.ClusterHealthIndicatorTypeClusterStatus,
			CurrentValue:       "HEALTH_OK",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeQuorum,
			CurrentValue:       "5 of 5",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
		{
			Indicator:          models.ClusterHealthIndicatorTypeOSDsDown,
			CurrentValue:       "0 of 15",
			CurrentValueStatus: models.ClusterHealthIndicatorStatusGood,
		},
	}, nil).Once()

	p.On("Green", "[%s] %s = %s", []any{
		"     GOOD", models.ClusterHealthIndicatorTypeClusterStatus, "HEALTH_OK",
	}).Return().Once()
	p.On("Green", "[%s] %s = %s", []any{
		"     GOOD", models.ClusterHealthIndicatorTypeQuorum, "5 of 5",
	}).Return().Once()
	p.On("Green", "[%s] %s = %s", []any{
		"     GOOD", models.ClusterHealthIndicatorTypeOSDsDown, "0 of 15",
	}).Return().Once()

	err := Healthcheck(context.Background(), HealthcheckConfig{
		Printer: p,
		Service: m,
	})
	r.NoError(err)
}

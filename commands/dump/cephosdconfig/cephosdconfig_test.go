package cephosdconfig

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
	"github.com/runityru/cephctl/printer"
	"github.com/runityru/cephctl/service"
)

func TestDumpCephOSDConfig(t *testing.T) {
	r := require.New(t)

	m := service.NewMock()
	defer m.AssertExpectations(t)

	p := printer.NewMock()
	defer p.AssertExpectations(t)

	m.On("DumpOSDConfig").Return(models.CephOSDConfig{
		AllowCrimson:           true,
		BackfillfullRatio:      0.9,
		FullRatio:              0.95,
		NearfullRatio:          0.85,
		RequireMinCompatClient: "reef",
	}, nil).Once()

	p.On("Println", []any{
		"kind: CephOSDConfig\nspec:\n    allow_crimson: true\n    backfillfull_ratio: 0.9\n    full_ratio: 0.95\n    nearfull_ratio: 0.85\n    require_min_compat_client: reef\n",
	}).Return().Once()

	err := DumpCephOSDConfig(context.Background(), DumpCephOSDConfigConfig{
		Printer: p,
		Service: m,
	})
	r.NoError(err)
}

package cephconfig

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
	"github.com/runityru/cephctl/printer"
	"github.com/runityru/cephctl/service"
)

func TestDumpCephConfig(t *testing.T) {
	r := require.New(t)

	m := service.NewMock()
	defer m.AssertExpectations(t)

	p := printer.NewMock()
	defer p.AssertExpectations(t)

	m.On("DumpConfig").Return(models.CephConfig{
		"global": {
			"key": "value",
		},
	}, nil).Once()

	p.On("Println", []any{"kind: CephConfig\nspec:\n    global:\n        key: value\n"}).Return().Once()

	err := DumpCephConfig(context.Background(), DumpCephConfigConfig{
		Printer: p,
		Service: m,
	})
	r.NoError(err)
}

package diff

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/go-ptr"

	"github.com/runityru/cephctl/models"
	"github.com/runityru/cephctl/printer"
	"github.com/runityru/cephctl/service"
)

func TestDiffCephConfig(t *testing.T) {
	r := require.New(t)

	m := service.NewMock()
	defer m.AssertExpectations(t)

	p := printer.NewMock()
	defer p.AssertExpectations(t)

	m.On("DiffCephConfig", models.CephConfig{
		"global": {
			"test": "value",
		},
	}).Return([]models.CephConfigDifference{
		{
			Kind:    models.CephConfigDifferenceKindAdd,
			Section: "mon",
			Key:     "test_key",
			Value:   ptr.String("value"),
		},
		{
			Kind:     models.CephConfigDifferenceKindChange,
			Section:  "osd.3",
			Key:      "test_key",
			OldValue: ptr.String("old_value"),
			Value:    ptr.String("value"),
		},
		{
			Kind:    models.CephConfigDifferenceKindRemove,
			Section: "osd",
			Key:     "test_key",
		},
	}, nil).Once()

	p.On("Green", "+ %s %s %s", []any{"mon", "test_key", "value"}).Return().Once()
	p.On("Yellow", "~ %s %s %s -> %s", []any{"osd.3", "test_key", "old_value", "value"}).Return().Once()
	p.On("Red", "- %s %s", []any{"osd", "test_key"}).Return().Once()

	err := Diff(context.Background(), DiffConfig{
		Printer:  p,
		Service:  m,
		SpecFile: "testdata/cephconfig.yaml",
	})
	r.NoError(err)
}

func TestDiffCephOSDConfig(t *testing.T) {
	r := require.New(t)

	m := service.NewMock()
	defer m.AssertExpectations(t)

	p := printer.NewMock()
	defer p.AssertExpectations(t)

	m.On("DiffCephOSDConfig", models.CephOSDConfig{
		AllowCrimson:           true,
		BackfillfullRatio:      0.9,
		FullRatio:              0.95,
		NearfullRatio:          0.85,
		RequireMinCompatClient: "luminous",
	}).Return([]models.CephOSDConfigDifference{
		{
			Key:      "allow_crimson",
			OldValue: "false",
			Value:    "true",
		},
	}, nil).Once()

	p.On("Yellow", "~ %s %s -> %s", []any{"allow_crimson", "false", "true"}).Return().Once()

	err := Diff(context.Background(), DiffConfig{
		Printer:  p,
		Service:  m,
		SpecFile: "testdata/cephosdconfig.yaml",
	})
	r.NoError(err)
}

package apply

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
	"github.com/runityru/cephctl/service"
)

func TestApplyCephConfig(t *testing.T) {
	r := require.New(t)

	m := service.NewMock()
	defer m.AssertExpectations(t)

	m.On("ApplyCephConfig", models.CephConfig{
		"global": {
			"test": "value",
		},
	}).Return(nil).Once()

	err := Apply(context.Background(), ApplyConfig{
		Service:  m,
		SpecFile: "testdata/cephconfig.yaml",
	})
	r.NoError(err)
}

func TestApplyCephOSDConfig(t *testing.T) {
	r := require.New(t)

	m := service.NewMock()
	defer m.AssertExpectations(t)

	m.On("ApplyCephOSDConfig", models.CephOSDConfig{
		AllowCrimson:           true,
		BackfillfullRatio:      0.9,
		FullRatio:              0.95,
		NearfullRatio:          0.85,
		RequireMinCompatClient: "luminous",
	}).Return(nil).Once()

	err := Apply(context.Background(), ApplyConfig{
		Service:  m,
		SpecFile: "testdata/cephosdconfig.yaml",
	})
	r.NoError(err)
}

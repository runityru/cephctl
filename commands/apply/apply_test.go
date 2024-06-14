package apply

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/service"
)

func TestApply(t *testing.T) {
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

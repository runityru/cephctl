package differ

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/teran/cephctl/models"
	ptr "github.com/teran/go-ptr"
)

func (s *differTestSuite) TestDiffCephConfig() {
	from := models.CephConfig{
		"osd": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "old_value",
		},
	}
	to := models.CephConfig{
		"mon": {
			"test_key": "value",
		},
		"osd.3": {
			"test_key": "value",
		},
	}

	diff, err := s.differ.DiffCephConfig(s.ctx, from, to)
	s.Require().NoError(err)
	s.Require().ElementsMatch([]models.CephConfigDifference{
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
	}, diff)
}

type differTestSuite struct {
	suite.Suite

	ctx    context.Context
	differ Differ
}

func (s *differTestSuite) SetupTest() {
	s.ctx = context.TODO()
	s.differ = New()
}

func (s *differTestSuite) TearDownTest() {
	s.ctx = nil
	s.differ = nil
}

func TestDifferTestSuite(t *testing.T) {
	suite.Run(t, &differTestSuite{})
}

func TestFlattenMap(t *testing.T) {
	r := require.New(t)

	out := flattenMap(map[string]map[string]string{
		"key1": {
			"key2": "value3",
		},
		"key3": {
			"key4": "value5",
		},
	})
	r.Equal(map[string]string{
		"key1:::key2": "value3",
		"key3:::key4": "value5",
	}, out)
}

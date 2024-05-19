package differ

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/teran/go-ptr"

	"github.com/teran/cephctl/models"
)

func (s *differTestSuite) TestDiffCephConfig() {
	type testCase struct {
		name     string
		from     models.CephConfig
		to       models.CephConfig
		expOut   []models.CephConfigDifference
		expError error
	}

	tcs := []testCase{
		{
			name: "ordinary config",
			from: models.CephConfig{
				"osd": {
					"test_key": "value",
				},
				"osd.3": {
					"test_key": "old_value",
				},
			},
			to: models.CephConfig{
				"mon": {
					"test_key": "value",
				},
				"osd.3": {
					"test_key": "value",
				},
			},
			expOut: []models.CephConfigDifference{
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
			},
		},
		{
			name:   "empty map",
			from:   models.CephConfig{},
			to:     models.CephConfig{},
			expOut: []models.CephConfigDifference{},
		},
		{
			name: "empty section",
			from: models.CephConfig{
				"": {
					"key": "value",
				},
			},
			to:       models.CephConfig{},
			expError: errors.Errorf("section name cannot be empty"),
		},
		{
			name: "empty key",
			from: models.CephConfig{
				"section": {
					"": "value",
				},
			},
			to:       models.CephConfig{},
			expError: errors.Errorf("key name cannot be empty"),
		},
	}

	for _, tc := range tcs {
		s.T().Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			diff, err := s.differ.DiffCephConfig(s.ctx, tc.from, tc.to)
			if tc.expError != nil {
				r.Error(err)
				r.Equal(tc.expError.Error(), err.Error())
			} else {
				r.NoError(err)
				r.NotNil(diff)
				r.ElementsMatch(tc.expOut, diff)
			}
		})
	}
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

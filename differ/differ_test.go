package differ

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/teran/go-ptr"

	"github.com/runityru/cephctl/models"
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

func (s *differTestSuite) TestDiffCephOSDConfig() {
	type testCase struct {
		name     string
		from     models.CephOSDConfig
		to       models.CephOSDConfig
		expOut   []models.CephOSDConfigDifference
		expError error
	}

	tcs := []testCase{
		{
			name: "ordinary config",
			from: models.CephOSDConfig{
				AllowCrimson:           false,
				NearfullRatio:          0.85,
				BackfillfullRatio:      0.9,
				FullRatio:              0.95,
				RequireMinCompatClient: "luminous",
			},
			to: models.CephOSDConfig{
				AllowCrimson:           true,
				NearfullRatio:          0.9,
				BackfillfullRatio:      0.95,
				FullRatio:              0.98,
				RequireMinCompatClient: "reef",
			},
			expOut: []models.CephOSDConfigDifference{
				{
					Key:      "allow_crimson",
					OldValue: "false",
					Value:    "true",
				},
				{
					Key:      "nearfull_ratio",
					OldValue: "0.85",
					Value:    "0.90",
				},
				{
					Key:      "backfillfull_ratio",
					OldValue: "0.90",
					Value:    "0.95",
				},
				{
					Key:      "full_ratio",
					OldValue: "0.95",
					Value:    "0.98",
				},
				{
					Key:      "require_min_compat_client",
					OldValue: "luminous",
					Value:    "reef",
				},
			},
		},
	}

	for _, tc := range tcs {
		s.T().Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			diff, err := s.differ.DiffCephOSDConfig(s.ctx, tc.from, tc.to)
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

// Definitions ...

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

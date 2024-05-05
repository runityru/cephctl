package models

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/teran/go-ptr"

	"github.com/teran/cephctl/models"
)

func TestStatusMapperEmpty(t *testing.T) {
	r := require.New(t)

	in := Status{}
	_, err := in.ToSvc()
	r.Error(err)
	r.Equal(errors.Cause(err), ErrUnexpectedInput)
}

func TestStatusMapper(t *testing.T) {
	r := require.New(t)

	in := Status{
		Health: StatusHealth{
			Status: "HEALTH_ERR",
			Checks: map[string]StatusCheck{
				"BLAH": {
					Severity: "HEALTH_ERR",
					Summary: StatusCheckSummary{
						Message: "some message",
						Count:   10,
					},
				},
			},
		},
		QuorumNames: []string{"host1", "host2", "host3", "host4", "host5"},
		MgrMap: StatusMgrMap{
			Available: true,
		},
		MonMap: StatusMonMap{
			NumMons: 5,
		},
		PGMap: StatusPGMap{
			NumPgs: 44,
			PgsByState: []PGsInState{
				{
					StateName: "active+degraded",
					Count:     10,
				},
				{
					StateName: "activating+clean",
					Count:     15,
				},
			},
		},
	}

	out, err := in.ToSvc()
	r.NoError(err)
	r.Equal(models.ClusterStatus{
		HealthStatus: models.ClusterStatusHealthERR,
		Checks: []models.ClusterStatusCheck{
			{
				Code:     "BLAH",
				Severity: models.ClusterStatusHealthERR,
				Summary:  "some message",
			},
		},
		MutedChecks:    []models.ClusterStatusMutedCheck{},
		QuorumAmount:   5,
		MonsTotal:      5,
		MonsDownAmount: 0,
		MGRsDownAmount: 0,
		MDSsDownAmount: 0,
		OSDsDownAmount: 0,
		UncleanPGs:     29,
		InactivePGs:    34,
	}, out)
}

func TestNewClusterStatusHealthFromString(t *testing.T) {
	type testCase struct {
		in             string
		expOut         models.ClusterStatusHealth
		expErrorString *string
		expErrorCause  error
	}

	tcs := []testCase{
		{
			in:     "HEALTH_OK",
			expOut: models.ClusterStatusHealthOK,
		},
		{
			in:     "HEALTH_WARN",
			expOut: models.ClusterStatusHealthWARN,
		},
		{
			in:     "HEALTH_ERR",
			expOut: models.ClusterStatusHealthERR,
		},
		{
			in:             "some_unexpected_health_status",
			expOut:         models.ClusterStatusHealthUnknown,
			expErrorString: ptr.String("some_unexpected_health_status: unexpected input"),
			expErrorCause:  ErrUnexpectedInput,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			r := require.New(t)

			ch, err := NewClusterStatusHealthFromString(tc.in)
			if tc.expErrorString != nil {
				r.Error(err)
				r.Equal(*tc.expErrorString, err.Error())
				r.EqualError(errors.Cause(err), tc.expErrorCause.Error())
			} else {
				r.NoError(err)
				r.Equal(tc.expOut, ch)
			}
		})
	}
}

func TestNewClusterStatusCheck(t *testing.T) {
	r := require.New(t)

	checks := map[string]StatusCheck{
		"BLAH": {
			Severity: "HEALTH_ERR",
			Summary: StatusCheckSummary{
				Message: "some message",
				Count:   10,
			},
		},
	}

	cs, err := NewClusterStatusCheck(checks)
	r.NoError(err)
	r.Equal([]models.ClusterStatusCheck{
		{
			Code:     "BLAH",
			Severity: "HEALTH_ERR",
			Summary:  "some message",
		},
	}, cs)
}

func TestNewClusterMutedChecks(t *testing.T) {
	r := require.New(t)

	mutes, err := NewClusterMutedChecks([]StatusHealthMute{
		{
			Code:    "BLAH",
			Sticky:  false,
			Summary: "test mute",
			Count:   4,
		},
	})
	r.NoError(err)
	r.Equal([]models.ClusterStatusMutedCheck{
		{
			Code:    "BLAH",
			Summary: "test mute",
		},
	}, mutes)
}

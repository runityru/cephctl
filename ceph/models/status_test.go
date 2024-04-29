package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
)

func TestStatusMapperEmpty(t *testing.T) {
	r := require.New(t)

	in := Status{}
	out, err := in.ToSvc()
	r.NoError(err)
	r.Equal(models.ClusterStatus{
		HealthStatus:   models.ClusterStatusHealthUnknown,
		Checks:         []models.ClusterStatusCheck{},
		MutedChecks:    []models.ClusterStatusMutedCheck{},
		QuorumAmount:   0,
		MonsTotal:      0,
		MonsDownAmount: 0,
		MGRsDownAmount: 1,
		MDSsDownAmount: 0,
		OSDsDownAmount: 0,
		UncleanPGs:     0,
		InactivePGs:    0,
	}, out)
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

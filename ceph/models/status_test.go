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
		HealthStatus: models.ClusterStatusHealthUnknown,
		Checks:       []models.ClusterStatusCheck{},
		MutedChecks:  []models.ClusterStatusMutedCheck{},
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
		MutedChecks: []models.ClusterStatusMutedCheck{},
	}, out)
}

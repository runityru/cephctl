package cluster_health

import (
	"github.com/teran/cephctl/models"
)

type testCase struct {
	name   string
	in     models.ClusterReport
	expOut models.ClusterHealthIndicator
}

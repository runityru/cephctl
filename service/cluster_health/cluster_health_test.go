package cluster_health

import (
	"github.com/runityru/cephctl/models"
)

type testCase struct {
	name   string
	in     models.ClusterReport
	expOut models.ClusterHealthIndicator
}

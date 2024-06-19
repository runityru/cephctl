package cluster_health

import (
	"context"

	"github.com/runityru/cephctl/models"
)

type ClusterHealthCheck func(ctx context.Context, cr models.ClusterReport) (models.ClusterHealthIndicator, error)

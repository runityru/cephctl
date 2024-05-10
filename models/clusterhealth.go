package models

type ClusterHealthIndicatorType string

const (
	// ClusterHealthIndicatorTypeClusterStatus reflects overall cluster status reported by ceph status command
	// Good: HEALTH_OK
	// AtRisk: HEALTH_WARN
	// Dangerous: HEALTH_ERR
	ClusterHealthIndicatorTypeClusterStatus ClusterHealthIndicatorType = "CLUSTER_STATUS"

	// ClusterHealthIndicatorTypeQuorum reflects monitor quorum status
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeQuorum ClusterHealthIndicatorType = "QUORUM"

	// ClusterHealthIndicatorTypeMonsDown reflects amount of monitor nodes which are down
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMonsDown ClusterHealthIndicatorType = "MON_DOWN"

	// ClusterHealthIndicatorTypeMgrsDown reflects amount of manager nodes which are down
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMgrsDown ClusterHealthIndicatorType = "MGR_DOWN"

	// ClusterHealthIndicatorTypeOSDsDown reflects amount of OSD nodes which are down
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeOSDsDown ClusterHealthIndicatorType = "OSD_DOWN"

	// ClusterHealthIndicatorTypeMDSsDown reflects amount of MDS nodes which are down
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMDSsDown ClusterHealthIndicatorType = "MDS_DOWN"

	// ClusterHealthIndicatorTypeMutesAmount reflects amount of mutes set on the cluster
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMutesAmount ClusterHealthIndicatorType = "MUTES_AMOUNT"

	// ClusterHealthIndicatorTypeUncleanPGs reflects amount of PGs which are not in clean state
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeUncleanPGs ClusterHealthIndicatorType = "UNCLEAN_PGS"

	// ClusterHealthIndicatorTypeInactivePGs reflects amount of PGs which are not in active state
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeInactivePGs ClusterHealthIndicatorType = "INACTIVE_PGS"

	// ClusterHealthIndicatorTypeAllowCrimson reflects allow_crimson flag state
	// Good: false
	// AtRisk: true
	// Dangerous: n/a
	ClusterHealthIndicatorTypeAllowCrimson ClusterHealthIndicatorType = "ALLOW_CRIMSON"
)

type ClusterHealthIndicatorStatus string

const (
	ClusterHealthIndicatorStatusGood      ClusterHealthIndicatorStatus = "GOOD"
	ClusterHealthIndicatorStatusAtRisk    ClusterHealthIndicatorStatus = "AT_RISK"
	ClusterHealthIndicatorStatusDangerous ClusterHealthIndicatorStatus = "DANGEROUS"
	ClusterHealthIndicatorStatusUnknown   ClusterHealthIndicatorStatus = "UNKNOWN"
)

type ClusterHealthIndicator struct {
	Indicator          ClusterHealthIndicatorType
	CurrentValue       string
	CurrentValueStatus ClusterHealthIndicatorStatus
}

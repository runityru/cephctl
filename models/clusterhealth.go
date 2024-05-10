package models

type ClusterHealthIndicatorType string

const (
	// ClusterHealthIndicatorTypeAllowCrimson reflects allow_crimson flag state
	// Good: false
	// AtRisk: true
	// Dangerous: n/a
	ClusterHealthIndicatorTypeAllowCrimson ClusterHealthIndicatorType = "ALLOW_CRIMSON"

	// ClusterHealthIndicatorTypeClusterStatus reflects overall cluster status reported by ceph status command
	// Good: HEALTH_OK
	// AtRisk: HEALTH_WARN
	// Dangerous: HEALTH_ERR
	ClusterHealthIndicatorTypeClusterStatus ClusterHealthIndicatorType = "CLUSTER_STATUS"

	// ClusterHealthIndicatorTypeInactivePGs reflects amount of PGs which are not in active state
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeInactivePGs ClusterHealthIndicatorType = "INACTIVE_PGS"

	// ClusterHealthIndicatorTypeMonsDown reflects amount of monitor nodes which are down
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMonsDown ClusterHealthIndicatorType = "MON_DOWN"

	// ClusterHealthIndicatorTypeMutesAmount reflects amount of mutes set on the cluster
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMutesAmount ClusterHealthIndicatorType = "MUTES_AMOUNT"

	// ClusterHealthIndicatorTypeOSDsDown reflects amount of OSD nodes which are down
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeOSDsDown ClusterHealthIndicatorType = "OSD_DOWN"

	// ClusterHealthIndicatorTypeOSDsMetadataSize reflects metadata size in percents of total OSD capacity
	// Good: 0-7
	// AtRisk: >7
	// Dangerous: >10
	ClusterHealthIndicatorTypeOSDsMetadataSize ClusterHealthIndicatorType = "OSD_METADATA_SIZE"

	// ClusterHealthIndicatorTypeQuorum reflects monitor quorum status
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeQuorum ClusterHealthIndicatorType = "QUORUM"

	// ClusterHealthIndicatorTypeUncleanPGs reflects amount of PGs which are not in clean state
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeUncleanPGs ClusterHealthIndicatorType = "UNCLEAN_PGS"
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

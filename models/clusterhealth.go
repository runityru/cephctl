package models

type ClusterHealthIndicatorType string

const (
	// ClusterHealthIndicatorTypeAllowCrimson reflects allow_crimson flag state
	//
	// Description: allow_crimson flag means possibility of using Crimson OSD
	// 	releases which are future releases so not marked as release or stable
	//  which makes it risky to use in production environment.
	//
	// Ref: https://docs.ceph.com/en/latest/glossary/#term-Crimson
	//
	// Good: false
	// AtRisk: true
	// Dangerous: n/a
	ClusterHealthIndicatorTypeAllowCrimson ClusterHealthIndicatorType = "ALLOW_CRIMSON"

	// ClusterHealthIndicatorTypeClusterStatus reflects overall cluster status
	// 	reported by ceph status command
	//
	// Description: cluster health is the universal indicator for overall cluster
	//	status which is also displayed via `ceph status` command.
	//
	// Ref: https://docs.ceph.com/en/latest/rados/operations/health-checks/
	//
	// Good: HEALTH_OK
	// AtRisk: HEALTH_WARN
	// Dangerous: HEALTH_ERR
	ClusterHealthIndicatorTypeClusterStatus ClusterHealthIndicatorType = "CLUSTER_STATUS"

	// ClusterHealthIndicatorTypeInactivePGs reflects amount of PGs which are not in
	// 	active state
	//
	// Description: Inactive PGs indicator shows how many PGs are inactive i.e. can not be
	// 	used to perform IO operations at the moment.
	//
	// Ref: https://docs.ceph.com/en/latest/rados/operations/monitoring-osd-pg/#monitoring-pg-states
	//
	// Good: 0
	// AtRisk: n/a
	// Dangerous: >0
	ClusterHealthIndicatorTypeInactivePGs ClusterHealthIndicatorType = "INACTIVE_PGS"

	// ClusterHealthIndicatorTypeMonsDown reflects amount of monitor nodes which are down
	//
	// Description: amount of monitors which are not up at the moment
	//
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMonsDown ClusterHealthIndicatorType = "MON_DOWN"

	// ClusterHealthIndicatorTypeMutesAmount reflects amount of mutes set on the cluster
	//
	// Description: Ceph allows to mute checks i.e. exclude them from triggering
	// 	overall cluster status. Muted checks are easy to miss when making decision
	//  of performing any maintenance which could cause a more serious cluster
	//  state or even data loss.
	//
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeMutesAmount ClusterHealthIndicatorType = "MUTES_AMOUNT"

	// ClusterHealthIndicatorTypeOSDsDown reflects amount of OSD nodes which are down
	//
	// Description: Amount of OSDs in down state
	//
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeOSDsDown ClusterHealthIndicatorType = "OSD_DOWN"

	// ClusterHealthIndicatorTypeOSDsMetadataSize reflects metadata size in percents of total OSD capacity
	//
	// Description: In Ceph documentation is described typical data usage for
	// 	block.db volume which is between 1-4%, however using RGW increases this
	//	estimation to 4% at least. This value allows to estimate size of block.db
	// 	volume for hybrid OSD and growing over block.db volume capacity will
	//  cause spillover - i.e. writing metadata to data volume which is usually
	//  much slower.
	//
	// Ref: https://docs.ceph.com/en/latest/rados/configuration/bluestore-config-ref/#sizing
	//
	// Good: 0-7
	// AtRisk: >15
	// Dangerous: >20
	ClusterHealthIndicatorTypeOSDsMetadataSize ClusterHealthIndicatorType = "OSD_METADATA_SIZE"

	// ClusterHealthIndicatorTypeOSDVersionMismatch indicates different versions
	// 	of running OSD daemons at the same time
	//
	// Description: running different versions of components is normal only
	// 	while upgrade procedure is a go. In all other cases daemon versions
	// 	should match i.e. their amount must be equal 1 except the case of
	// 	upgrade.
	//
	// Good: 1
	// AtRisk: 2
	// Dangerous: >2
	ClusterHealthIndicatorTypeOSDsNumDaemonVersions ClusterHealthIndicatorType = "OSD_NUM_DAEMON_VERSIONS"

	// ClusterHealthIndicatorTypeOSDsOut reflects amount of OSD nodes which are out
	//
	// Description: Amount of OSDs in out state
	//
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeOSDsOut ClusterHealthIndicatorType = "OSD_OUT"

	// ClusterHealthIndicatorTypeQuorum reflects monitor quorum status
	//
	// Description: monitors in quorum which should be the same as total
	// 	monitors amount
	//
	// Good: 0
	// AtRisk: >0
	// Dangerous: n/a
	ClusterHealthIndicatorTypeQuorum ClusterHealthIndicatorType = "QUORUM"

	// ClusterHealthIndicatorTypeUncleanPGs reflects amount of PGs which are not in clean state
	//
	// Description: Inactive PGs indicator shows how many PGs are inactive i.e. can not be
	// 	used to perform IO operations at the moment.
	//
	// Ref: https://docs.ceph.com/en/latest/rados/operations/monitoring-osd-pg/#monitoring-pg-states
	//
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

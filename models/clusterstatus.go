package models

type ClusterStatusHealth string

const (
	ClusterStatusHealthOK      ClusterStatusHealth = "HEALTH_OK"
	ClusterStatusHealthWARN    ClusterStatusHealth = "HEALTH_WARN"
	ClusterStatusHealthERR     ClusterStatusHealth = "HEALTH_ERR"
	ClusterStatusHealthUnknown ClusterStatusHealth = "UNKNOWN"
)

type ClusterStatusMutedCheck struct {
	Code    string
	Summary string
}

type ClusterStatusCheck struct {
	Code     string
	Severity ClusterStatusHealth
	Summary  string
}

type ClusterStatus struct {
	HealthStatus   ClusterStatusHealth
	Checks         []ClusterStatusCheck
	MutedChecks    []ClusterStatusMutedCheck
	QuorumAmount   uint
	MonsTotal      uint
	MonsDownAmount uint
	MGRsDownAmount uint
	MDSsDownAmount uint
	OSDsDownAmount uint
	UncleanPGs     uint
	InactivePGs    uint
}

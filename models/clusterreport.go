package models

type ClusterReport struct {
	HealthStatus                 ClusterStatusHealth
	Checks                       []ClusterStatusCheck
	MutedChecks                  []ClusterStatusMutedCheck
	NumMons                      uint8
	NumMonsInQuorum              uint8
	AllowCrimson                 bool
	StretchMode                  bool
	NumOSDs                      uint16
	NumOSDsWithoutClusterAddress uint16
	NumOSDsIn                    uint16
	NumOSDsUp                    uint16
	NumOSDsByRelease             map[string]uint16
	NumOSDsByVersion             map[string]uint16
	NumOSDsByDeviceType          map[string]uint16
	Devices                      []Device
	TotalOSDCapacityKB           uint64
	TotalOSDUsedDataKB           uint64
	TotalOSDUsedMetaKB           uint64
	TotalOSDUsedOMAPKB           uint64
	NumPools                     uint16
	NumPGs                       uint32
	NumPGsByState                map[string]uint32
}

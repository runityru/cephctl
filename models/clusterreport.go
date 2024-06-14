package models

type ClusterReport struct {
	AllowCrimson                 bool
	Checks                       []ClusterStatusCheck
	Devices                      []Device
	HealthStatus                 ClusterStatusHealth
	MutedChecks                  []ClusterStatusMutedCheck
	NumMons                      uint8
	NumMonsInQuorum              uint8
	NumOSDs                      uint16
	NumOSDsByDeviceType          map[string]uint16
	NumOSDsByRelease             map[string]uint16
	NumOSDsByVersion             map[string]uint16
	NumOSDsIn                    uint16
	NumOSDsUp                    uint16
	NumOSDsWithoutClusterAddress uint16
	NumPGs                       uint32
	NumPGsByState                map[string]uint32
	NumPools                     uint16
	StretchMode                  bool
	TotalOSDCapacityKB           uint64
	TotalOSDUsedDataKB           uint64
	TotalOSDUsedMetaKB           uint64
	TotalOSDUsedOMAPKB           uint64
}

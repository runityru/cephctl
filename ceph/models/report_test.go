package models

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/runityru/cephctl/models"
)

func TestReportServiceMapServicesRgwDaemonGenericMapUnmarshalJSON(t *testing.T) {
	r := require.New(t)

	data, err := os.ReadFile("testdata/ReportServiceMapServicesRgwDaemonGenericMap.json")
	r.NoError(err)

	gm := ReportServiceMapServicesRgw{}
	err = json.Unmarshal(data, &gm)
	r.NoError(err)
	r.Equal(ReportServiceMapServicesRgw{
		Daemons: ReportServiceMapServicesRgwDaemonGenericMap{
			"84305071": ReportServiceMapServicesRgwDaemon{
				StartEpoch: 163222,
				StartStamp: "2024-04-27T06:04:24.624576+0000",
				Gid:        84305071,
				Addr:       "192.168.1.81:0/692915515",
				Metadata: ReportServiceMapServicesRgwDaemonMetadata{
					Arch:              "x86_64",
					CephRelease:       "reef",
					CephVersion:       "ceph version 18.2.2 (531c0d11a1c5d39fbfe6aa8a521f023abf3bf3e2) reef (stable)",
					CephVersionShort:  "18.2.2",
					CPU:               "12th Gen Intel(R) Core(TM) i5-1240P",
					Distro:            "centos",
					DistroDescription: "CentOS Stream 8",
					DistroVersion:     "8",
					FrontendConfig0:   "beast port=7480",
					FrontendType0:     "beast",
					Hostname:          "radosgw-bmrq7",
					ID:                "radosgw.rgw01",
					KernelDescription: "#1 SMP PREEMPT_DYNAMIC Thu Apr 4 22:31:43 UTC 2024",
					KernelVersion:     "5.14.0-362.24.1.el9_3.0.1.x86_64",
					MemSwapKb:         "0",
					MemTotalKb:        "16114028",
					NumHandles:        "1",
					OS:                "Linux",
					PID:               "1",
					ZoneID:            "b35fb671-d4e5-4132-b9f9-c411add2f9e7",
					ZoneName:          "default",
					ZonegroupID:       "2c7837d2-74c6-4166-a703-f101281119cc",
					ZonegroupName:     "default",
				},
			},
		},
	}, gm)
}

func TestReportToSvc(t *testing.T) {
	type testCase struct {
		name   string
		in     string
		expOut models.ClusterReport
	}

	osdDaemons := []models.OSDDaemon{
		{
			ID:               0,
			Hostname:         "nuc01",
			Architecture:     "x86_64",
			FrontIP:          "192.168.1.201",
			BackIP:           "192.168.2.231",
			MemoryTotalBytes: 65414276,
			SwapTotalBytes:   0,
			IsRotational:     false,
			Devices: []string{
				"sda",
			},
		},
		{
			ID:               1,
			Hostname:         "nuc01",
			Architecture:     "x86_64",
			FrontIP:          "192.168.1.201",
			BackIP:           "192.168.2.231",
			MemoryTotalBytes: 65414276,
			SwapTotalBytes:   0,
			IsRotational:     false,
			Devices: []string{
				"nvme0n1",
			},
		},
	}

	tcs := []testCase{
		{
			name: "clean cluster",
			in:   "testdata/report_samples/CleanReport.json",
			expOut: models.ClusterReport{
				HealthStatus:    models.ClusterStatusHealthOK,
				Checks:          []models.ClusterStatusCheck{},
				MutedChecks:     []models.ClusterStatusMutedCheck{},
				NumMons:         5,
				NumMonsInQuorum: 5,
				NumOSDs:         15,
				NumOSDsIn:       15,
				NumOSDsUp:       15,
				NumOSDsByRelease: map[string]uint16{
					"reef": 2,
				},
				NumOSDsByVersion: map[string]uint16{
					"18.2.2": 2,
				},
				NumOSDsByDeviceType: map[string]uint16{
					"ssd": 2,
				},
				TotalOSDCapacityKB: uint64(22_321_704_960),
				TotalOSDUsedDataKB: uint64(10_986_978_208),
				TotalOSDUsedMetaKB: uint64(512_967_627),
				TotalOSDUsedOMAPKB: uint64(5_822_580),
				NumPools:           14,
				NumPGs:             330,
				NumPGsByState: map[string]uint32{
					"active":        330,
					"backfill_wait": 50,
					"backfilling":   2,
					"clean":         278,
					"remapped":      52,
				},
				OSDDaemons:             osdDaemons,
				AllowCrimson:           false,
				NearfullRatio:          0.85,
				BackfillfullRatio:      0.9,
				FullRatio:              0.95,
				RequireMinCompatClient: "luminous",
			},
		},
		{
			name: "cluster with OSDs in out state",
			in:   "testdata/report_samples/ReportWithOutOSDs.json",
			expOut: models.ClusterReport{
				HealthStatus: models.ClusterStatusHealthWARN,
				Checks: []models.ClusterStatusCheck{
					{
						Code:     "OSDMAP_FLAGS",
						Severity: "HEALTH_WARN",
						Summary:  "nodown,noout flag(s) set",
					},
					{
						Code:     "PG_BACKFILL_FULL",
						Severity: "HEALTH_WARN",
						Summary:  "Low space hindering backfill (add storage if this doesn't resolve itself): 14 pgs backfill_toofull",
					},
				},
				MutedChecks:     []models.ClusterStatusMutedCheck{},
				NumMons:         5,
				NumMonsInQuorum: 5,
				NumOSDs:         15,
				NumOSDsIn:       13,
				NumOSDsUp:       15,
				NumOSDsByRelease: map[string]uint16{
					"reef": 2,
				},
				NumOSDsByVersion: map[string]uint16{
					"18.2.2": 2,
				},
				NumOSDsByDeviceType: map[string]uint16{
					"ssd": 2,
				},
				TotalOSDCapacityKB: 18_729_263_104,
				TotalOSDUsedDataKB: 10_393_147_824,
				TotalOSDUsedMetaKB: 489_119_531,
				TotalOSDUsedOMAPKB: 1_478_996,
				NumPools:           14,
				NumPGs:             330,
				NumPGsByState: map[string]uint32{
					"active":           330,
					"backfill_toofull": 14,
					"backfill_wait":    78,
					"backfilling":      2,
					"clean":            250,
					"remapped":         153,
				},
				OSDDaemons:             osdDaemons,
				AllowCrimson:           false,
				NearfullRatio:          0.85,
				BackfillfullRatio:      0.9,
				FullRatio:              0.95,
				RequireMinCompatClient: "luminous",
			},
		},
		{
			name: "cluster with OSDs in down state",
			in:   "testdata/report_samples/ReportWithDownOSDs.json",
			expOut: models.ClusterReport{
				HealthStatus: models.ClusterStatusHealthWARN,
				Checks: []models.ClusterStatusCheck{
					{
						Code:     "FS_DEGRADED",
						Severity: "HEALTH_WARN",
						Summary:  "1 filesystem is degraded",
					},
					{
						Code:     "MDS_SLOW_METADATA_IO",
						Severity: "HEALTH_WARN",
						Summary:  "1 MDSs report slow metadata IOs",
					},
					{
						Code:     "OSD_DOWN",
						Severity: "HEALTH_WARN",
						Summary:  "6 osds down",
					},
					{
						Code:     "OSD_HOST_DOWN",
						Severity: "HEALTH_WARN",
						Summary:  "2 hosts (6 osds) down",
					},
					{
						Code:     "OSD_SLOW_PING_TIME_BACK",
						Severity: "HEALTH_WARN",
						Summary:  "Slow OSD heartbeats on back (longest 83761.606ms)",
					},
					{
						Code:     "OSD_SLOW_PING_TIME_FRONT",
						Severity: "HEALTH_WARN",
						Summary:  "Slow OSD heartbeats on front (longest 84721.090ms)",
					},
					{
						Code:     "PG_AVAILABILITY",
						Severity: "HEALTH_WARN",
						Summary:  "Reduced data availability: 212 pgs inactive, 192 pgs down",
					},
					{
						Code:     "PG_DEGRADED",
						Severity: "HEALTH_WARN",
						Summary:  "Degraded data redundancy: 13932/28946801 objects degraded (0.048%), 105 pgs degraded, 111 pgs undersized",
					},
				},
				MutedChecks:     []models.ClusterStatusMutedCheck{},
				NumMons:         5,
				NumMonsInQuorum: 5,
				NumOSDs:         14,
				NumOSDsIn:       14,
				NumOSDsUp:       8,
				NumOSDsByRelease: map[string]uint16{
					"reef": 2,
				},
				NumOSDsByVersion: map[string]uint16{
					"18.2.2": 2,
				},
				NumOSDsByDeviceType: map[string]uint16{
					"ssd": 2,
				},
				TotalOSDCapacityKB: 16_985_456_640,
				TotalOSDUsedDataKB: 7_161_124_160,
				TotalOSDUsedMetaKB: 2_194_240_219,
				TotalOSDUsedOMAPKB: 1_172_260,
				NumPools:           14,
				NumPGs:             330,
				NumPGsByState: map[string]uint32{
					"active":     118,
					"clean":      27,
					"degraded":   105,
					"down":       192,
					"peered":     20,
					"undersized": 111,
				},
				OSDDaemons:             osdDaemons,
				AllowCrimson:           false,
				NearfullRatio:          0.85,
				BackfillfullRatio:      0.9,
				FullRatio:              0.95,
				RequireMinCompatClient: "luminous",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			data, err := os.ReadFile(tc.in)
			r.NoError(err)

			rep := &Report{}
			err = json.Unmarshal(data, rep)
			r.NoError(err)

			out, err := rep.ToSvc()
			r.NoError(err)
			r.Equal(tc.expOut, out)
		})
	}
}

func TestCountOSDs(t *testing.T) {
	r := require.New(t)

	total, up, in, withoutClusterAddress := countOSDs([]ReportOSDMapOSD{
		{In: 1},
		{In: 1},
		{Up: 1},
		{In: 1, Up: 1, ClusterAddrs: ReportOSDMapOSDClusterAddrs{
			Addrvec: []ReportOSDMapOSDClusterAddrsAddrvec{
				{}, {},
			},
		}},
	})
	r.Equal(uint16(4), total)
	r.Equal(uint16(2), up)
	r.Equal(uint16(3), in)
	r.Equal(uint16(3), withoutClusterAddress)
}

func TestCountOSDsByRelease(t *testing.T) {
	r := require.New(t)

	c := countOSDsByRelease([]ReportOSDMetadata{})
	r.Equal(map[string]uint16{}, c)

	c = countOSDsByRelease([]ReportOSDMetadata{
		{CephRelease: "quincy"},
		{CephRelease: "quincy"},
		{CephRelease: "reef"},
		{CephRelease: "reef"},
		{CephRelease: "reef"},
	})
	r.Equal(map[string]uint16{
		"quincy": 2,
		"reef":   3,
	}, c)
}

func TestCountOSDsByVersion(t *testing.T) {
	r := require.New(t)

	c := countOSDsByVersion([]ReportOSDMetadata{})
	r.Equal(map[string]uint16{}, c)

	c = countOSDsByVersion([]ReportOSDMetadata{
		{CephVersionShort: "17.2.6"},
		{CephVersionShort: "17.2.6"},
		{CephVersionShort: "18.2.1"},
		{CephVersionShort: "18.2.2"},
	})
	r.Equal(map[string]uint16{
		"17.2.6": 2,
		"18.2.1": 1,
		"18.2.2": 1,
	}, c)
}

func TestCountOSDsByDeviceType(t *testing.T) {
	r := require.New(t)

	c := countOSDsByDeviceType([]ReportOSDMetadata{})
	r.Equal(map[string]uint16{}, c)

	c = countOSDsByDeviceType([]ReportOSDMetadata{
		{BluestoreBdevType: "ssd"},
		{BluestoreBdevType: "hdd"},
		{BluestoreBdevType: "ssd"},
		{BluestoreBdevType: "ssd"},
	})
	r.Equal(map[string]uint16{
		"ssd": 3,
		"hdd": 1,
	}, c)
}

func TestCountPGs(t *testing.T) {
	r := require.New(t)

	total, byState, err := countPGs([]ReportNumPGByState{
		{
			State: "active+remapped+backfill_wait",
			Num:   2,
		},
		{
			State: "active+remapped+backfilling",
			Num:   4,
		},
		{
			State: "active+clean",
			Num:   8,
		},
		{
			State: "down",
			Num:   10,
		},
	})
	r.NoError(err)
	r.Equal(uint32(24), total)
	r.Equal(map[string]uint32{
		"active":        14,
		"remapped":      6,
		"backfill_wait": 2,
		"backfilling":   4,
		"clean":         8,
		"down":          10,
	}, byState)
}

func TestParseCephIPAddress(t *testing.T) {
	type testCase struct {
		name     string
		in       string
		expOut   string
		expError error
	}

	tcs := []testCase{
		{
			name:   "valid ceph address string",
			in:     "[v2:192.168.2.232:6802/2418,v1:192.168.2.232:6803/2418]",
			expOut: "192.168.2.232",
		},
		{
			name:     "invalid ceph address string",
			in:       "blah",
			expError: errors.New("malformed ceph address string"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			out, err := parseCephIPAddress(tc.in)
			if tc.expError != nil {
				r.Error(err)
				r.Equal(tc.expError.Error(), err.Error())
			} else {
				r.NoError(err)
				r.Equal(tc.expOut, out)
			}
		})
	}
}

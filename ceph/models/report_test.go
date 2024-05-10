package models

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teran/cephctl/models"
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
	r := require.New(t)

	data, err := os.ReadFile("testdata/Report.json")
	r.NoError(err)

	rep := &Report{}
	err = json.Unmarshal(data, rep)
	r.NoError(err)

	out, err := rep.ToSvc()
	r.NoError(err)
	r.Equal(models.ClusterReport{
		HealthStatus:    models.ClusterStatusHealthOK,
		Checks:          []models.ClusterStatusCheck{},
		MutedChecks:     []models.ClusterStatusMutedCheck{},
		NumMons:         5,
		NumMonsInQuorum: 5,
		NumOSDs:         15,
		NumOSDsIn:       15,
		NumOSDsUp:       15,
		NumOSDsByRelease: map[string]uint16{
			"reef": 15,
		},
		NumOSDsByVersion: map[string]uint16{
			"18.2.2": 15,
		},
		NumOSDsByDeviceType: map[string]uint16{
			"ssd": 15,
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
	}, out)
}

func TestCountOSDs(t *testing.T) {
	r := require.New(t)

	up, in, withoutClusterAddress := countOSDs([]ReportOSDMapOSD{
		{In: 1},
		{In: 1},
		{Up: 1},
		{In: 1, Up: 1, ClusterAddrs: ReportOSDMapOSDClusterAddrs{
			Addrvec: []ReportOSDMapOSDClusterAddrsAddrvec{
				{}, {},
			},
		}},
	})
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
	})
	r.NoError(err)
	r.Equal(uint32(14), total)
	r.Equal(map[string]uint32{
		"active":        14,
		"remapped":      6,
		"backfill_wait": 2,
		"backfilling":   4,
		"clean":         8,
	}, byState)
}

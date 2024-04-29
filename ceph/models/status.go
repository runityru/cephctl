package models

import (
	"strings"

	"github.com/teran/cephctl/models"
)

type StatusCheckSummary struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

type StatusCheck struct {
	Severity string             `json:"severity"`
	Summary  StatusCheckSummary `json:"summary"`
	Muted    bool               `json:"muted"`
}

type StatusHealthMute struct {
	Code    string `json:"code"`
	Sticky  bool   `json:"sticky"`
	Summary string `json:"summary"`
	Count   int    `json:"count"`
}

type StatusHealth struct {
	Status string                 `json:"status"`
	Checks map[string]StatusCheck `json:"checks"`
	Mutes  []StatusHealthMute     `json:"mutes"`
}

type StatusMonMap struct {
	Epoch             int    `json:"epoch"`
	MinMonReleaseName string `json:"min_mon_release_name"`
	NumMons           int    `json:"num_mons"`
}

type StatusOSDMap struct {
	Epoch          int `json:"epoch"`
	NumOsds        int `json:"num_osds"`
	NumUpOsds      int `json:"num_up_osds"`
	OsdUpSince     int `json:"osd_up_since"`
	NumInOsds      int `json:"num_in_osds"`
	OsdInSince     int `json:"osd_in_since"`
	NumRemappedPgs int `json:"num_remapped_pgs"`
}

type PGsInState struct {
	StateName string `json:"state_name"`
	Count     int    `json:"count"`
}

type StatusPGMap struct {
	PgsByState    []PGsInState `json:"pgs_by_state"`
	NumPgs        int          `json:"num_pgs"`
	NumPools      int          `json:"num_pools"`
	NumObjects    int          `json:"num_objects"`
	DataBytes     int64        `json:"data_bytes"`
	BytesUsed     int64        `json:"bytes_used"`
	BytesAvail    int64        `json:"bytes_avail"`
	BytesTotal    int64        `json:"bytes_total"`
	ReadBytesSec  int          `json:"read_bytes_sec"`
	WriteBytesSec int          `json:"write_bytes_sec"`
	ReadOpPerSec  int          `json:"read_op_per_sec"`
	WriteOpPerSec int          `json:"write_op_per_sec"`
}

type StatusFSMapByRank struct {
	FilesystemID int    `json:"filesystem_id"`
	Rank         int    `json:"rank"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	Gid          int    `json:"gid"`
}

type StatusFSMap struct {
	Epoch     int                 `json:"epoch"`
	ID        int                 `json:"id"`
	Up        int                 `json:"up"`
	In        int                 `json:"in"`
	Max       int                 `json:"max"`
	ByRank    []StatusFSMapByRank `json:"by_rank"`
	UpStandby int                 `json:"up:standby"`
}

type StatusMgrMap struct {
	Available   bool     `json:"available"`
	NumStandbys int      `json:"num_standbys"`
	Modules     []string `json:"modules"`
	Services    struct {
		Prometheus string `json:"prometheus"`
	} `json:"services"`
}

type Status struct {
	FSID          string       `json:"fsid"`
	Health        StatusHealth `json:"health"`
	ElectionEpoch int          `json:"election_epoch"`
	Quorum        []int        `json:"quorum"`
	QuorumNames   []string     `json:"quorum_names"`
	QuorumAge     int          `json:"quorum_age"`
	MonMap        StatusMonMap `json:"monmap"`
	OSDMap        StatusOSDMap `json:"osdmap"`
	PGMap         StatusPGMap  `json:"pgmap"`
	FSMap         StatusFSMap  `json:"fsmap"`
	MgrMap        StatusMgrMap `json:"mgrmap"`
}

func (st *Status) ToSvc() (models.ClusterStatus, error) {
	csh := models.ClusterStatusHealthUnknown
	switch st.Health.Status {
	case "HEALTH_OK":
		csh = models.ClusterStatusHealthOK
	case "HEALTH_WARN":
		csh = models.ClusterStatusHealthWARN
	case "HEALTH_ERR":
		csh = models.ClusterStatusHealthERR
	}

	checks := []models.ClusterStatusCheck{}
	for code, c := range st.Health.Checks {
		severity := models.ClusterStatusHealthUnknown
		switch c.Severity {
		case "HEALTH_OK":
			severity = models.ClusterStatusHealthOK
		case "HEALTH_WARN":
			severity = models.ClusterStatusHealthWARN
		case "HEALTH_ERR":
			severity = models.ClusterStatusHealthERR
		}

		checks = append(checks, models.ClusterStatusCheck{
			Code:     code,
			Severity: severity,
			Summary:  c.Summary.Message,
		})
	}

	mutes := []models.ClusterStatusMutedCheck{}
	for _, m := range st.Health.Mutes {
		mutes = append(mutes, models.ClusterStatusMutedCheck{
			Code:    m.Code,
			Summary: m.Summary,
		})
	}

	monsDown := st.MonMap.NumMons - len(st.QuorumNames)
	osdsDown := st.OSDMap.NumOsds - st.OSDMap.NumUpOsds

	var mgrsDown uint = 0
	if !st.MgrMap.Available {
		mgrsDown++
	}

	mdsDown := st.FSMap.Max - st.FSMap.Up

	pgStates := map[string]uint{
		"clean":  0,
		"active": 0,
	}
	for _, pg := range st.PGMap.PgsByState {
		for _, state := range strings.Split(pg.StateName, "+") {
			if _, ok := pgStates[state]; !ok {
				pgStates[state] = 0
			}

			pgStates[state] += uint(pg.Count)
		}
	}

	return models.ClusterStatus{
		HealthStatus:   csh,
		Checks:         checks,
		MutedChecks:    mutes,
		MonsTotal:      uint(st.MonMap.NumMons),
		QuorumAmount:   uint(len(st.QuorumNames)),
		MonsDownAmount: uint(monsDown),
		MGRsDownAmount: mgrsDown,
		MDSsDownAmount: uint(mdsDown),
		OSDsDownAmount: uint(osdsDown),
		UncleanPGs:     uint(st.PGMap.NumPgs) - pgStates["clean"],
		InactivePGs:    uint(st.PGMap.NumPgs) - pgStates["active"],
	}, nil
}

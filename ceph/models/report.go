package models

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/runityru/cephctl/models"
)

var ErrOverflow = errors.Errorf("unexpected overflow")

type ReportHealth struct {
	Status string                 `json:"status"`
	Checks map[string]StatusCheck `json:"checks"`
	Mutes  []StatusHealthMute     `json:"mutes"`
}

type ReportMonMapMon struct {
	Rank        int    `json:"rank"`
	Name        string `json:"name"`
	PublicAddrs struct {
		Addrvec []struct {
			Type  string `json:"type"`
			Addr  string `json:"addr"`
			Nonce int    `json:"nonce"`
		} `json:"addrvec"`
	} `json:"public_addrs"`
	Addr          string  `json:"addr"`
	PublicAddr    string  `json:"public_addr"`
	Priority      int     `json:"priority"`
	Weight        float64 `json:"weight"`
	CrushLocation string  `json:"crush_location"`
}

type ReportMonMap struct {
	Epoch             int       `json:"epoch"`
	Fsid              string    `json:"fsid"`
	Modified          time.Time `json:"modified"`
	Created           time.Time `json:"created"`
	MinMonRelease     int       `json:"min_mon_release"`
	MinMonReleaseName string    `json:"min_mon_release_name"`
	ElectionStrategy  int       `json:"election_strategy"`
	DisallowedLeaders string    `json:"disallowed_leaders: "`
	StretchMode       bool      `json:"stretch_mode"`
	TiebreakerMon     string    `json:"tiebreaker_mon"`
	RemovedRanks      string    `json:"removed_ranks: "`
	Features          struct {
		Persistent []string `json:"persistent"`
		Optional   []any    `json:"optional"`
	} `json:"features"`
	Mons []ReportMonMapMon `json:"mons"`
}

type ReportOSDMapPoolLastPGMergeMeta struct {
	SourcePGID       string `json:"source_pgid"`
	ReadyEpoch       int    `json:"ready_epoch"`
	LastEpochStarted int    `json:"last_epoch_started"`
	LastEpochClean   int    `json:"last_epoch_clean"`
	SourceVersion    string `json:"source_version"`
	TargetVersion    string `json:"target_version"`
}

type ReportOSDMapPoolHitSetParams struct {
	Type string `json:"type"`
}

type ReportOSDMapPoolOptions struct {
	PGAutoscaleBias  int `json:"pg_autoscale_bias"`
	PGNumMax         int `json:"pg_num_max"`
	PGNumMin         int `json:"pg_num_min"`
	RecoveryPriority int `json:"recovery_priority"`
}

type ReportOSDMapPoolApplicationMetadata struct {
	Mgr    struct{} `json:"mgr"`
	RBD    struct{} `json:"rbd"`
	CephFS struct {
		Data string `json:"data"`
	} `json:"cephfs"`
	Rgw struct{} `json:"rgw"`
}

type ReportOSDMapPoolReadBalance struct {
	ScoreActing                    float64 `json:"score_acting"`
	ScoreStable                    float64 `json:"score_stable"`
	OptimalScore                   float64 `json:"optimal_score"`
	RawScoreActing                 float64 `json:"raw_score_acting"`
	RawScoreStable                 float64 `json:"raw_score_stable"`
	PrimaryAffinityWeighted        float64 `json:"primary_affinity_weighted"`
	AveragePrimaryAffinity         float64 `json:"average_primary_affinity"`
	AveragePrimaryAffinityWeighted float64 `json:"average_primary_affinity_weighted"`
}

type ReportOSDMapPool struct {
	Pool                              int                                 `json:"pool"`
	PoolName                          string                              `json:"pool_name"`
	CreateTime                        string                              `json:"create_time"`
	Flags                             int                                 `json:"flags"`
	FlagsNames                        string                              `json:"flags_names"`
	Type                              int                                 `json:"type"`
	Size                              int                                 `json:"size"`
	MinSize                           int                                 `json:"min_size"`
	CrushRule                         int                                 `json:"crush_rule"`
	PeeringCrushBucketCount           int                                 `json:"peering_crush_bucket_count"`
	PeeringCrushBucketTarget          int                                 `json:"peering_crush_bucket_target"`
	PeeringCrushBucketBarrier         int                                 `json:"peering_crush_bucket_barrier"`
	PeeringCrushBucketMandatoryMember int64                               `json:"peering_crush_bucket_mandatory_member"`
	ObjectHash                        int                                 `json:"object_hash"`
	PgAutoscaleMode                   string                              `json:"pg_autoscale_mode"`
	PgNum                             int                                 `json:"pg_num"`
	PgPlacementNum                    int                                 `json:"pg_placement_num"`
	PgPlacementNumTarget              int                                 `json:"pg_placement_num_target"`
	PgNumTarget                       int                                 `json:"pg_num_target"`
	PgNumPending                      int                                 `json:"pg_num_pending"`
	LastPgMergeMeta                   ReportOSDMapPoolLastPGMergeMeta     `json:"last_pg_merge_meta"`
	LastChange                        string                              `json:"last_change"`
	LastForceOpResend                 string                              `json:"last_force_op_resend"`
	LastForceOpResendPrenautilus      string                              `json:"last_force_op_resend_prenautilus"`
	LastForceOpResendPreluminous      string                              `json:"last_force_op_resend_preluminous"`
	Auid                              int                                 `json:"auid"`
	SnapMode                          string                              `json:"snap_mode"`
	SnapSeq                           int                                 `json:"snap_seq"`
	SnapEpoch                         int                                 `json:"snap_epoch"`
	PoolSnaps                         []any                               `json:"pool_snaps"`
	RemovedSnaps                      string                              `json:"removed_snaps"`
	QuotaMaxBytes                     int                                 `json:"quota_max_bytes"`
	QuotaMaxObjects                   int                                 `json:"quota_max_objects"`
	Tiers                             []any                               `json:"tiers"`
	TierOf                            int                                 `json:"tier_of"`
	ReadTier                          int                                 `json:"read_tier"`
	WriteTier                         int                                 `json:"write_tier"`
	CacheMode                         string                              `json:"cache_mode"`
	TargetMaxBytes                    int                                 `json:"target_max_bytes"`
	TargetMaxObjects                  int                                 `json:"target_max_objects"`
	CacheTargetDirtyRatioMicro        int                                 `json:"cache_target_dirty_ratio_micro"`
	CacheTargetDirtyHighRatioMicro    int                                 `json:"cache_target_dirty_high_ratio_micro"`
	CacheTargetFullRatioMicro         int                                 `json:"cache_target_full_ratio_micro"`
	CacheMinFlushAge                  int                                 `json:"cache_min_flush_age"`
	CacheMinEvictAge                  int                                 `json:"cache_min_evict_age"`
	ErasureCodeProfile                string                              `json:"erasure_code_profile"`
	HitSetParams                      ReportOSDMapPoolHitSetParams        `json:"hit_set_params"`
	HitSetPeriod                      int                                 `json:"hit_set_period"`
	HitSetCount                       int                                 `json:"hit_set_count"`
	UseGmtHitset                      bool                                `json:"use_gmt_hitset"`
	MinReadRecencyForPromote          int                                 `json:"min_read_recency_for_promote"`
	MinWriteRecencyForPromote         int                                 `json:"min_write_recency_for_promote"`
	HitSetGradeDecayRate              int                                 `json:"hit_set_grade_decay_rate"`
	HitSetSearchLastN                 int                                 `json:"hit_set_search_last_n"`
	GradeTable                        []any                               `json:"grade_table"`
	StripeWidth                       int                                 `json:"stripe_width"`
	ExpectedNumObjects                int                                 `json:"expected_num_objects"`
	FastRead                          bool                                `json:"fast_read"`
	Options                           ReportOSDMapPoolOptions             `json:"options,omitempty"`
	ApplicationMetadata               ReportOSDMapPoolApplicationMetadata `json:"application_metadata,omitempty"`
	ReadBalance                       ReportOSDMapPoolReadBalance         `json:"read_balance,omitempty"`
}

type ReportOSDMapOSDPublicAddrs struct {
	Addrvec []struct {
		Type  string `json:"type"`
		Addr  string `json:"addr"`
		Nonce int    `json:"nonce"`
	} `json:"addrvec"`
}

type ReportOSDMapOSDClusterAddrsAddrvec struct {
	Type  string `json:"type"`
	Addr  string `json:"addr"`
	Nonce int    `json:"nonce"`
}

type ReportOSDMapOSDClusterAddrs struct {
	Addrvec []ReportOSDMapOSDClusterAddrsAddrvec `json:"addrvec"`
}

type ReportOSDMapOSDHeartbeatBackAddrs struct {
	Addrvec []struct {
		Type  string `json:"type"`
		Addr  string `json:"addr"`
		Nonce int    `json:"nonce"`
	} `json:"addrvec"`
}

type ReportOSDMapOSDHeartbeatFrontAddrs struct {
	Addrvec []struct {
		Type  string `json:"type"`
		Addr  string `json:"addr"`
		Nonce int    `json:"nonce"`
	} `json:"addrvec"`
}

type ReportOSDMapOSD struct {
	Osd                 int                                `json:"osd"`
	UUID                string                             `json:"uuid"`
	Up                  uint8                              `json:"up"`
	In                  uint8                              `json:"in"`
	Weight              float64                            `json:"weight"`
	PrimaryAffinity     int                                `json:"primary_affinity"`
	LastCleanBegin      int                                `json:"last_clean_begin"`
	LastCleanEnd        int                                `json:"last_clean_end"`
	UpFrom              int                                `json:"up_from"`
	UpThru              int                                `json:"up_thru"`
	DownAt              int                                `json:"down_at"`
	LostAt              int                                `json:"lost_at"`
	PublicAddrs         ReportOSDMapOSDPublicAddrs         `json:"public_addrs"`
	ClusterAddrs        ReportOSDMapOSDClusterAddrs        `json:"cluster_addrs"`
	HeartbeatBackAddrs  ReportOSDMapOSDHeartbeatBackAddrs  `json:"heartbeat_back_addrs"`
	HeartbeatFrontAddrs ReportOSDMapOSDHeartbeatFrontAddrs `json:"heartbeat_front_addrs"`
	PublicAddr          string                             `json:"public_addr"`
	ClusterAddr         string                             `json:"cluster_addr"`
	HeartbeatBackAddr   string                             `json:"heartbeat_back_addr"`
	HeartbeatFrontAddr  string                             `json:"heartbeat_front_addr"`
	State               []string                           `json:"state"`
}

type ReportOSDMapOSDXInfo struct {
	OSD                  int     `json:"osd"`
	DownStamp            string  `json:"down_stamp"`
	LaggyProbability     float64 `json:"laggy_probability"`
	LaggyInterval        int     `json:"laggy_interval"`
	Features             int64   `json:"features"`
	OldWeight            float64 `json:"old_weight"`
	LastPurgedSnapsScrub string  `json:"last_purged_snaps_scrub"`
	DeadEpoch            int     `json:"dead_epoch"`
}

type ReportOSDMapErasureCodeProfile struct {
	CrushDeviceClass          string `json:"crush-device-class"`
	CrushFailureDomain        string `json:"crush-failure-domain"`
	CrushRoot                 string `json:"crush-root"`
	Directory                 string `json:"directory"`
	JerasurePerChunkAlignment string `json:"jerasure-per-chunk-alignment"`
	K                         string `json:"k"`
	M                         string `json:"m"`
	Packetsize                string `json:"packetsize"`
	Plugin                    string `json:"plugin"`
	Technique                 string `json:"technique"`
	W                         string `json:"w"`
}

type ReportOSDMapStretchMode struct {
	StretchModeEnabled    bool `json:"stretch_mode_enabled"`
	StretchBucketCount    int  `json:"stretch_bucket_count"`
	DegradedStretchMode   int  `json:"degraded_stretch_mode"`
	RecoveringStretchMode int  `json:"recovering_stretch_mode"`
	StretchModeBucket     int  `json:"stretch_mode_bucket"`
}

type ReportOSDMap struct {
	Epoch                  int                    `json:"epoch"`
	Fsid                   string                 `json:"fsid"`
	Created                string                 `json:"created"`
	Modified               string                 `json:"modified"`
	LastUpChange           string                 `json:"last_up_change"`
	LastInChange           string                 `json:"last_in_change"`
	Flags                  string                 `json:"flags"`
	FlagsNum               int                    `json:"flags_num"`
	FlagsSet               []string               `json:"flags_set"`
	CrushVersion           int                    `json:"crush_version"`
	FullRatio              float32                `json:"full_ratio"`
	BackfillfullRatio      float32                `json:"backfillfull_ratio"`
	NearfullRatio          float32                `json:"nearfull_ratio"`
	ClusterSnapshot        string                 `json:"cluster_snapshot"`
	PoolMax                int                    `json:"pool_max"`
	MaxOsd                 int                    `json:"max_osd"`
	RequireMinCompatClient string                 `json:"require_min_compat_client"`
	MinCompatClient        string                 `json:"min_compat_client"`
	RequireOsdRelease      string                 `json:"require_osd_release"`
	AllowCrimson           bool                   `json:"allow_crimson"`
	Pools                  []ReportOSDMapPool     `json:"pools"`
	OSDs                   []ReportOSDMapOSD      `json:"osds"`
	OSDXInfo               []ReportOSDMapOSDXInfo `json:"osd_xinfo"`
	PGUpmap                []any                  `json:"pg_upmap"`
	PGUpmapItems           []struct {
		Pgid     string `json:"pgid"`
		Mappings []struct {
			From int `json:"from"`
			To   int `json:"to"`
		} `json:"mappings"`
	} `json:"pg_upmap_items"`
	PGUpmapPrimaries    []any                                     `json:"pg_upmap_primaries"`
	PGTemp              []any                                     `json:"pg_temp"`
	PrimaryTemp         []any                                     `json:"primary_temp"`
	Blocklist           struct{}                                  `json:"blocklist"`
	RangeBlocklist      struct{}                                  `json:"range_blocklist"`
	ErasureCodeProfiles map[string]ReportOSDMapErasureCodeProfile `json:"erasure_code_profiles"`
	RemovedSnapsQueue   []any                                     `json:"removed_snaps_queue"`
	NewRemovedSnaps     []any                                     `json:"new_removed_snaps"`
	NewPurgedSnaps      []any                                     `json:"new_purged_snaps"`
	CrushNodeFlags      struct{}                                  `json:"crush_node_flags"`
	DeviceClassFlags    struct{}                                  `json:"device_class_flags"`
	StretchMode         ReportOSDMapStretchMode                   `json:"stretch_mode"`
}

type ReportOSDMetadata struct {
	ID                            int       `json:"id"`
	Arch                          string    `json:"arch"`
	BackAddr                      string    `json:"back_addr"`
	BackIface                     string    `json:"back_iface"`
	Bluefs                        string    `json:"bluefs"`
	BluefsDedicatedDb             string    `json:"bluefs_dedicated_db"`
	BluefsDedicatedWal            string    `json:"bluefs_dedicated_wal"`
	BluefsSingleSharedDevice      string    `json:"bluefs_single_shared_device"`
	BluestoreBdevAccessMode       string    `json:"bluestore_bdev_access_mode"`
	BluestoreBdevBlockSize        string    `json:"bluestore_bdev_block_size"`
	BluestoreBdevDevNode          string    `json:"bluestore_bdev_dev_node"`
	BluestoreBdevDevices          string    `json:"bluestore_bdev_devices"`
	BluestoreBdevDriver           string    `json:"bluestore_bdev_driver"`
	BluestoreBdevOptimalIoSize    string    `json:"bluestore_bdev_optimal_io_size"`
	BluestoreBdevPartitionPath    string    `json:"bluestore_bdev_partition_path"`
	BluestoreBdevRotational       string    `json:"bluestore_bdev_rotational"`
	BluestoreBdevSize             string    `json:"bluestore_bdev_size"`
	BluestoreBdevSupportDiscard   string    `json:"bluestore_bdev_support_discard"`
	BluestoreBdevType             string    `json:"bluestore_bdev_type"`
	BluestoreMinAllocSize         string    `json:"bluestore_min_alloc_size"`
	CephRelease                   string    `json:"ceph_release"`
	CephVersion                   string    `json:"ceph_version"`
	CephVersionShort              string    `json:"ceph_version_short"`
	CephVersionWhenCreated        string    `json:"ceph_version_when_created"`
	CPU                           string    `json:"cpu"`
	CreatedAt                     time.Time `json:"created_at"`
	DefaultDeviceClass            string    `json:"default_device_class"`
	DeviceIds                     string    `json:"device_ids"`
	DevicePaths                   string    `json:"device_paths"`
	Devices                       string    `json:"devices"`
	Distro                        string    `json:"distro"`
	DistroDescription             string    `json:"distro_description"`
	DistroVersion                 string    `json:"distro_version"`
	FrontAddr                     string    `json:"front_addr"`
	FrontIface                    string    `json:"front_iface"`
	HbBackAddr                    string    `json:"hb_back_addr"`
	HbFrontAddr                   string    `json:"hb_front_addr"`
	Hostname                      string    `json:"hostname"`
	JournalRotational             string    `json:"journal_rotational"`
	KernelDescription             string    `json:"kernel_description"`
	KernelVersion                 string    `json:"kernel_version"`
	MemSwapKb                     string    `json:"mem_swap_kb"`
	MemTotalKb                    string    `json:"mem_total_kb"`
	NetworkNumaUnknownIfaces      string    `json:"network_numa_unknown_ifaces"`
	ObjectstoreNumaUnknownDevices string    `json:"objectstore_numa_unknown_devices,omitempty"`
	Os                            string    `json:"os"`
	OsdData                       string    `json:"osd_data"`
	OsdObjectstore                string    `json:"osd_objectstore"`
	OsdspecAffinity               string    `json:"osdspec_affinity"`
	Rotational                    string    `json:"rotational"`
	ObjectstoreNumaNode           string    `json:"objectstore_numa_node,omitempty"`
	ObjectstoreNumaNodes          string    `json:"objectstore_numa_nodes,omitempty"`
}

type ReportCRUSHMapTunables struct {
	ChooseLocalTries         int    `json:"choose_local_tries"`
	ChooseLocalFallbackTries int    `json:"choose_local_fallback_tries"`
	ChooseTotalTries         int    `json:"choose_total_tries"`
	ChooseleafDescendOnce    int    `json:"chooseleaf_descend_once"`
	ChooseleafVaryR          int    `json:"chooseleaf_vary_r"`
	ChooseleafStable         int    `json:"chooseleaf_stable"`
	StrawCalcVersion         int    `json:"straw_calc_version"`
	AllowedBucketAlgs        int    `json:"allowed_bucket_algs"`
	Profile                  string `json:"profile"`
	OptimalTunables          int    `json:"optimal_tunables"`
	LegacyTunables           int    `json:"legacy_tunables"`
	MinimumRequiredVersion   string `json:"minimum_required_version"`
	RequireFeatureTunables   int    `json:"require_feature_tunables"`
	RequireFeatureTunables2  int    `json:"require_feature_tunables2"`
	HasV2Rules               int    `json:"has_v2_rules"`
	RequireFeatureTunables3  int    `json:"require_feature_tunables3"`
	HasV3Rules               int    `json:"has_v3_rules"`
	HasV4Buckets             int    `json:"has_v4_buckets"`
	RequireFeatureTunables5  int    `json:"require_feature_tunables5"`
	HasV5Rules               int    `json:"has_v5_rules"`
}

type ReportCRUSHMapRuleStep struct {
	Op       string `json:"op"`
	Item     int    `json:"item,omitempty"`
	ItemName string `json:"item_name,omitempty"`
	Num      int    `json:"num,omitempty"`
	Type     string `json:"type,omitempty"`
}

type ReportCRUSHMapRule struct {
	RuleID   int                      `json:"rule_id"`
	RuleName string                   `json:"rule_name"`
	Type     int                      `json:"type"`
	Steps    []ReportCRUSHMapRuleStep `json:"steps"`
}

type ReportCRUSHMapBucketItem struct {
	ID     int     `json:"id"`
	Weight float64 `json:"weight"`
	Pos    int     `json:"pos"`
}

type ReportCRUSHMapBucket struct {
	ID       int                        `json:"id"`
	Name     string                     `json:"name"`
	TypeID   int                        `json:"type_id"`
	TypeName string                     `json:"type_name"`
	Weight   float64                    `json:"weight"`
	Alg      string                     `json:"alg"`
	Hash     string                     `json:"hash"`
	Items    []ReportCRUSHMapBucketItem `json:"items"`
}

type ReportCRUSHMapType struct {
	TypeID int    `json:"type_id"`
	Name   string `json:"name"`
}

type ReportCRUSHMapDevice struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

type ReportCRUSHMap struct {
	Devices    []ReportCRUSHMapDevice `json:"devices"`
	Types      []ReportCRUSHMapType   `json:"types"`
	Buckets    []ReportCRUSHMapBucket `json:"buckets"`
	Rules      []ReportCRUSHMapRule   `json:"rules"`
	Tunables   ReportCRUSHMapTunables `json:"tunables"`
	ChooseArgs struct {
		Num1 []struct {
			BucketID  int         `json:"bucket_id"`
			WeightSet [][]float64 `json:"weight_set"`
		} `json:"-1"`
	} `json:"choose_args"`
}

type ReportOSDMapCleanEpochs struct {
	MinLastEpochClean int `json:"min_last_epoch_clean"`
	LastEpochClean    struct {
		PerPool []struct {
			Poolid int `json:"poolid"`
			Floor  int `json:"floor"`
		} `json:"per_pool"`
	} `json:"last_epoch_clean"`
	OSDEpochs []struct {
		ID    int `json:"id"`
		Epoch int `json:"epoch"`
	} `json:"osd_epochs"`
}

type ReportFSMap struct {
	Epoch        int `json:"epoch"`
	DefaultFscid int `json:"default_fscid"`
	Compat       struct {
		Compat   struct{} `json:"compat"`
		RoCompat struct{} `json:"ro_compat"`
		Incompat struct {
			Feature1  string `json:"feature_1"`
			Feature2  string `json:"feature_2"`
			Feature3  string `json:"feature_3"`
			Feature4  string `json:"feature_4"`
			Feature5  string `json:"feature_5"`
			Feature6  string `json:"feature_6"`
			Feature8  string `json:"feature_8"`
			Feature9  string `json:"feature_9"`
			Feature10 string `json:"feature_10"`
		} `json:"incompat"`
	} `json:"compat"`
	FeatureFlags struct {
		EnableMultiple      bool `json:"enable_multiple"`
		EverEnabledMultiple bool `json:"ever_enabled_multiple"`
	} `json:"feature_flags"`
	Standbys []struct {
		Gid         int    `json:"gid"`
		Name        string `json:"name"`
		Rank        int    `json:"rank"`
		Incarnation int    `json:"incarnation"`
		State       string `json:"state"`
		StateSeq    int    `json:"state_seq"`
		Addr        string `json:"addr"`
		Addrs       struct {
			Addrvec []struct {
				Type  string `json:"type"`
				Addr  string `json:"addr"`
				Nonce int    `json:"nonce"`
			} `json:"addrvec"`
		} `json:"addrs"`
		JoinFscid     int   `json:"join_fscid"`
		ExportTargets []any `json:"export_targets"`
		Features      int64 `json:"features"`
		Flags         int   `json:"flags"`
		Compat        struct {
			Compat   struct{} `json:"compat"`
			RoCompat struct{} `json:"ro_compat"`
			Incompat struct {
				Feature1  string `json:"feature_1"`
				Feature2  string `json:"feature_2"`
				Feature3  string `json:"feature_3"`
				Feature4  string `json:"feature_4"`
				Feature5  string `json:"feature_5"`
				Feature6  string `json:"feature_6"`
				Feature7  string `json:"feature_7"`
				Feature8  string `json:"feature_8"`
				Feature9  string `json:"feature_9"`
				Feature10 string `json:"feature_10"`
			} `json:"incompat"`
		} `json:"compat"`
		Epoch int `json:"epoch"`
	} `json:"standbys"`
	Filesystems []struct {
		Mdsmap struct {
			Epoch      int `json:"epoch"`
			Flags      int `json:"flags"`
			FlagsState struct {
				Joinable            bool `json:"joinable"`
				AllowSnaps          bool `json:"allow_snaps"`
				AllowMultimdsSnaps  bool `json:"allow_multimds_snaps"`
				AllowStandbyReplay  bool `json:"allow_standby_replay"`
				RefuseClientSession bool `json:"refuse_client_session"`
			} `json:"flags_state"`
			EverAllowedFeatures       int      `json:"ever_allowed_features"`
			ExplicitlyAllowedFeatures int      `json:"explicitly_allowed_features"`
			Created                   string   `json:"created"`
			Modified                  string   `json:"modified"`
			Tableserver               int      `json:"tableserver"`
			Root                      int      `json:"root"`
			SessionTimeout            int      `json:"session_timeout"`
			SessionAutoclose          int      `json:"session_autoclose"`
			RequiredClientFeatures    struct{} `json:"required_client_features"`
			MaxFileSize               int64    `json:"max_file_size"`
			LastFailure               int      `json:"last_failure"`
			LastFailureOsdEpoch       int      `json:"last_failure_osd_epoch"`
			Compat                    struct {
				Compat   struct{} `json:"compat"`
				RoCompat struct{} `json:"ro_compat"`
				Incompat struct {
					Feature1  string `json:"feature_1"`
					Feature2  string `json:"feature_2"`
					Feature3  string `json:"feature_3"`
					Feature4  string `json:"feature_4"`
					Feature5  string `json:"feature_5"`
					Feature6  string `json:"feature_6"`
					Feature7  string `json:"feature_7"`
					Feature8  string `json:"feature_8"`
					Feature9  string `json:"feature_9"`
					Feature10 string `json:"feature_10"`
				} `json:"incompat"`
			} `json:"compat"`
			MaxMds int   `json:"max_mds"`
			In     []int `json:"in"`
			Up     struct {
				Mds0 int `json:"mds_0"`
			} `json:"up"`
			Failed  []any `json:"failed"`
			Damaged []any `json:"damaged"`
			Stopped []int `json:"stopped"`
			Info    struct {
				Gid83261581 struct {
					Gid         int    `json:"gid"`
					Name        string `json:"name"`
					Rank        int    `json:"rank"`
					Incarnation int    `json:"incarnation"`
					State       string `json:"state"`
					StateSeq    int    `json:"state_seq"`
					Addr        string `json:"addr"`
					Addrs       struct {
						Addrvec []struct {
							Type  string `json:"type"`
							Addr  string `json:"addr"`
							Nonce int    `json:"nonce"`
						} `json:"addrvec"`
					} `json:"addrs"`
					JoinFscid     int   `json:"join_fscid"`
					ExportTargets []any `json:"export_targets"`
					Features      int64 `json:"features"`
					Flags         int   `json:"flags"`
					Compat        struct {
						Compat   struct{} `json:"compat"`
						RoCompat struct{} `json:"ro_compat"`
						Incompat struct {
							Feature1  string `json:"feature_1"`
							Feature2  string `json:"feature_2"`
							Feature3  string `json:"feature_3"`
							Feature4  string `json:"feature_4"`
							Feature5  string `json:"feature_5"`
							Feature6  string `json:"feature_6"`
							Feature7  string `json:"feature_7"`
							Feature8  string `json:"feature_8"`
							Feature9  string `json:"feature_9"`
							Feature10 string `json:"feature_10"`
						} `json:"incompat"`
					} `json:"compat"`
				} `json:"gid_83261581"`
			} `json:"info"`
			DataPools          []int  `json:"data_pools"`
			MetadataPool       int    `json:"metadata_pool"`
			Enabled            bool   `json:"enabled"`
			FsName             string `json:"fs_name"`
			Balancer           string `json:"balancer"`
			BalRankMask        string `json:"bal_rank_mask"`
			StandbyCountWanted int    `json:"standby_count_wanted"`
		} `json:"mdsmap"`
		ID int `json:"id"`
	} `json:"filesystems"`
}

type ReportAuth struct {
	FirstCommitted int `json:"first_committed"`
	LastCommitted  int `json:"last_committed"`
	NumSecrets     int `json:"num_secrets"`
}

type ReportPoolSumStatSum struct {
	NumBytes                   int64 `json:"num_bytes"`
	NumObjects                 int   `json:"num_objects"`
	NumObjectClones            int   `json:"num_object_clones"`
	NumObjectCopies            int   `json:"num_object_copies"`
	NumObjectsMissingOnPrimary int   `json:"num_objects_missing_on_primary"`
	NumObjectsMissing          int   `json:"num_objects_missing"`
	NumObjectsDegraded         int   `json:"num_objects_degraded"`
	NumObjectsMisplaced        int   `json:"num_objects_misplaced"`
	NumObjectsUnfound          int   `json:"num_objects_unfound"`
	NumObjectsDirty            int   `json:"num_objects_dirty"`
	NumWhiteouts               int   `json:"num_whiteouts"`
	NumRead                    int   `json:"num_read"`
	NumReadKb                  int64 `json:"num_read_kb"`
	NumWrite                   int   `json:"num_write"`
	NumWriteKb                 int64 `json:"num_write_kb"`
	NumScrubErrors             int   `json:"num_scrub_errors"`
	NumShallowScrubErrors      int   `json:"num_shallow_scrub_errors"`
	NumDeepScrubErrors         int   `json:"num_deep_scrub_errors"`
	NumObjectsRecovered        int   `json:"num_objects_recovered"`
	NumBytesRecovered          int64 `json:"num_bytes_recovered"`
	NumKeysRecovered           int   `json:"num_keys_recovered"`
	NumObjectsOmap             int   `json:"num_objects_omap"`
	NumObjectsHitSetArchive    int   `json:"num_objects_hit_set_archive"`
	NumBytesHitSetArchive      int   `json:"num_bytes_hit_set_archive"`
	NumFlush                   int   `json:"num_flush"`
	NumFlushKb                 int   `json:"num_flush_kb"`
	NumEvict                   int   `json:"num_evict"`
	NumEvictKb                 int   `json:"num_evict_kb"`
	NumPromote                 int   `json:"num_promote"`
	NumFlushModeHigh           int   `json:"num_flush_mode_high"`
	NumFlushModeLow            int   `json:"num_flush_mode_low"`
	NumEvictModeSome           int   `json:"num_evict_mode_some"`
	NumEvictModeFull           int   `json:"num_evict_mode_full"`
	NumObjectsPinned           int   `json:"num_objects_pinned"`
	NumLegacySnapsets          int   `json:"num_legacy_snapsets"`
	NumLargeOmapObjects        int   `json:"num_large_omap_objects"`
	NumObjectsManifest         int   `json:"num_objects_manifest"`
	NumOmapBytes               int   `json:"num_omap_bytes"`
	NumOmapKeys                int   `json:"num_omap_keys"`
	NumObjectsRepaired         int   `json:"num_objects_repaired"`
}

type ReportPoolSumStoreStats struct {
	Total                   int `json:"total"`
	Available               int `json:"available"`
	InternallyReserved      int `json:"internally_reserved"`
	Allocated               int `json:"allocated"`
	DataStored              int `json:"data_stored"`
	DataCompressed          int `json:"data_compressed"`
	DataCompressedAllocated int `json:"data_compressed_allocated"`
	DataCompressedOriginal  int `json:"data_compressed_original"`
	OmapAllocated           int `json:"omap_allocated"`
	InternalMetadata        int `json:"internal_metadata"`
}

type ReportPoolSum struct {
	StatSum       ReportPoolSumStatSum    `json:"stat_sum"`
	StoreStats    ReportPoolSumStoreStats `json:"store_stats"`
	LogSize       int                     `json:"log_size"`
	OndiskLogSize int                     `json:"ondisk_log_size"`
	Up            int                     `json:"up"`
	Acting        int                     `json:"acting"`
	NumStoreStats int                     `json:"num_store_stats"`
}

type ReportOSDSumStatfs struct {
	Total                   int64 `json:"total"`
	Available               int64 `json:"available"`
	InternallyReserved      int64 `json:"internally_reserved"`
	Allocated               int64 `json:"allocated"`
	DataStored              int64 `json:"data_stored"`
	DataCompressed          int64 `json:"data_compressed"`
	DataCompressedAllocated int64 `json:"data_compressed_allocated"`
	DataCompressedOriginal  int64 `json:"data_compressed_original"`
	OmapAllocated           int64 `json:"omap_allocated"`
	InternalMetadata        int64 `json:"internal_metadata"`
}

type ReportOSDSumOpQueueAgeHist struct {
	Histogram  []any `json:"histogram"`
	UpperBound int   `json:"upper_bound"`
}

type ReportOSDSumPerfStat struct {
	CommitLatencyMs int `json:"commit_latency_ms"`
	ApplyLatencyMs  int `json:"apply_latency_ms"`
	CommitLatencyNs int `json:"commit_latency_ns"`
	ApplyLatencyNs  int `json:"apply_latency_ns"`
}

type ReportOSDSumClass struct {
	UpFrom             int                        `json:"up_from"`
	Seq                int                        `json:"seq"`
	NumPGs             uint32                     `json:"num_pgs"`
	NumOSDs            uint16                     `json:"num_osds"`
	NumPerPoolOSDs     uint16                     `json:"num_per_pool_osds"`
	NumPerPoolOMAPOSDs uint16                     `json:"num_per_pool_omap_osds"`
	Kb                 uint64                     `json:"kb"`
	KbUsed             uint64                     `json:"kb_used"`
	KbUsedData         uint64                     `json:"kb_used_data"`
	KbUsedOmap         uint64                     `json:"kb_used_omap"`
	KbUsedMeta         uint64                     `json:"kb_used_meta"`
	KbAvail            uint64                     `json:"kb_avail"`
	Statfs             ReportOSDSumStatfs         `json:"statfs"`
	HbPeers            []any                      `json:"hb_peers"`
	SnapTrimQueueLen   uint                       `json:"snap_trim_queue_len"`
	NumSnapTrimming    int                        `json:"num_snap_trimming"`
	NumShardsRepaired  int                        `json:"num_shards_repaired"`
	OpQueueAgeHist     ReportOSDSumOpQueueAgeHist `json:"op_queue_age_hist"`
	PerfStat           ReportOSDSumPerfStat       `json:"perf_stat"`
	Alerts             []any                      `json:"alerts"`
	NetworkPingTimes   []any                      `json:"network_ping_times"`
}

type ReportOSDSum struct {
	UpFrom             int                `json:"up_from"`
	Seq                int                `json:"seq"`
	NumPGs             int                `json:"num_pgs"`
	NumOSDs            uint16             `json:"num_osds"`
	NumPerPoolOSDs     uint16             `json:"num_per_pool_osds"`
	NumPerPoolOMAPOSDs uint16             `json:"num_per_pool_omap_osds"`
	Kb                 uint64             `json:"kb"`
	KbUsed             uint64             `json:"kb_used"`
	KbUsedData         uint64             `json:"kb_used_data"`
	KbUsedOmap         uint64             `json:"kb_used_omap"`
	KbUsedMeta         uint64             `json:"kb_used_meta"`
	KbAvail            uint64             `json:"kb_avail"`
	Statfs             ReportOSDSumStatfs `json:"statfs"`
	HbPeers            []any              `json:"hb_peers"`
	SnapTrimQueueLen   int                `json:"snap_trim_queue_len"`
	NumSnapTrimming    int                `json:"num_snap_trimming"`
	NumShardsRepaired  int                `json:"num_shards_repaired"`
	OpQueueAgeHist     struct {
		Histogram  []any `json:"histogram"`
		UpperBound int   `json:"upper_bound"`
	} `json:"op_queue_age_hist"`
	PerfStat struct {
		CommitLatencyMs int `json:"commit_latency_ms"`
		ApplyLatencyMs  int `json:"apply_latency_ms"`
		CommitLatencyNs int `json:"commit_latency_ns"`
		ApplyLatencyNs  int `json:"apply_latency_ns"`
	} `json:"perf_stat"`
	Alerts           []any `json:"alerts"`
	NetworkPingTimes []any `json:"network_ping_times"`
}

type ReportPoolStats struct {
	Poolid  int    `json:"poolid"`
	NumPG   uint32 `json:"num_pg"`
	StatSum struct {
		NumBytes                   int64 `json:"num_bytes"`
		NumObjects                 int   `json:"num_objects"`
		NumObjectClones            int   `json:"num_object_clones"`
		NumObjectCopies            int   `json:"num_object_copies"`
		NumObjectsMissingOnPrimary int   `json:"num_objects_missing_on_primary"`
		NumObjectsMissing          int   `json:"num_objects_missing"`
		NumObjectsDegraded         int   `json:"num_objects_degraded"`
		NumObjectsMisplaced        int   `json:"num_objects_misplaced"`
		NumObjectsUnfound          int   `json:"num_objects_unfound"`
		NumObjectsDirty            int   `json:"num_objects_dirty"`
		NumWhiteouts               int   `json:"num_whiteouts"`
		NumRead                    int   `json:"num_read"`
		NumReadKb                  int64 `json:"num_read_kb"`
		NumWrite                   int   `json:"num_write"`
		NumWriteKb                 int64 `json:"num_write_kb"`
		NumScrubErrors             int   `json:"num_scrub_errors"`
		NumShallowScrubErrors      int   `json:"num_shallow_scrub_errors"`
		NumDeepScrubErrors         int   `json:"num_deep_scrub_errors"`
		NumObjectsRecovered        int   `json:"num_objects_recovered"`
		NumBytesRecovered          int64 `json:"num_bytes_recovered"`
		NumKeysRecovered           int   `json:"num_keys_recovered"`
		NumObjectsOmap             int   `json:"num_objects_omap"`
		NumObjectsHitSetArchive    int   `json:"num_objects_hit_set_archive"`
		NumBytesHitSetArchive      int   `json:"num_bytes_hit_set_archive"`
		NumFlush                   int   `json:"num_flush"`
		NumFlushKb                 int   `json:"num_flush_kb"`
		NumEvict                   int   `json:"num_evict"`
		NumEvictKb                 int   `json:"num_evict_kb"`
		NumPromote                 int   `json:"num_promote"`
		NumFlushModeHigh           int   `json:"num_flush_mode_high"`
		NumFlushModeLow            int   `json:"num_flush_mode_low"`
		NumEvictModeSome           int   `json:"num_evict_mode_some"`
		NumEvictModeFull           int   `json:"num_evict_mode_full"`
		NumObjectsPinned           int   `json:"num_objects_pinned"`
		NumLegacySnapsets          int   `json:"num_legacy_snapsets"`
		NumLargeOmapObjects        int   `json:"num_large_omap_objects"`
		NumObjectsManifest         int   `json:"num_objects_manifest"`
		NumOmapBytes               int   `json:"num_omap_bytes"`
		NumOmapKeys                int   `json:"num_omap_keys"`
		NumObjectsRepaired         int   `json:"num_objects_repaired"`
	} `json:"stat_sum"`
	StoreStats struct {
		Total                   int   `json:"total"`
		Available               int   `json:"available"`
		InternallyReserved      int   `json:"internally_reserved"`
		Allocated               int64 `json:"allocated"`
		DataStored              int64 `json:"data_stored"`
		DataCompressed          int   `json:"data_compressed"`
		DataCompressedAllocated int   `json:"data_compressed_allocated"`
		DataCompressedOriginal  int   `json:"data_compressed_original"`
		OmapAllocated           int   `json:"omap_allocated"`
		InternalMetadata        int   `json:"internal_metadata"`
	} `json:"store_stats"`
	LogSize       int `json:"log_size"`
	OndiskLogSize int `json:"ondisk_log_size"`
	Up            int `json:"up"`
	Acting        int `json:"acting"`
	NumStoreStats int `json:"num_store_stats"`
}

type ReportPaxos struct {
	FirstCommitted int `json:"first_committed"`
	LastCommitted  int `json:"last_committed"`
	LastPn         int `json:"last_pn"`
	AcceptedPn     int `json:"accepted_pn"`
}

type ReportOSDStats struct {
	OSD int   `json:"osd"`
	Seq int64 `json:"seq"`
}

type ReportNumPGByState struct {
	State string `json:"state"`
	Num   uint32 `json:"num"`
}

type ReportNumPGByOSD struct {
	OSD              int `json:"osd"`
	NumPrimaryPG     int `json:"num_primary_pg"`
	NumActingPG      int `json:"num_acting_pg"`
	NumUpNotActingPG int `json:"num_up_not_acting_pg"`
}

type ReportPurgedSnaps struct {
	Pool        int      `json:"pool"`
	PurgedSnaps struct{} `json:"purged_snaps"`
}

type ReportServiceMapServicesOSD struct {
	Daemons struct {
		Num11 struct {
			StartEpoch int      `json:"start_epoch"`
			StartStamp string   `json:"start_stamp"`
			Gid        int      `json:"gid"`
			Addr       string   `json:"addr"`
			Metadata   struct{} `json:"metadata"`
			TaskStatus struct{} `json:"task_status"`
		} `json:"11"`
		Summary string `json:"summary"`
	} `json:"daemons"`
}

type ReportServiceMapServicesRgwDaemonMetadata struct {
	Arch              string `json:"arch"`
	CephRelease       string `json:"ceph_release"`
	CephVersion       string `json:"ceph_version"`
	CephVersionShort  string `json:"ceph_version_short"`
	CPU               string `json:"cpu"`
	Distro            string `json:"distro"`
	DistroDescription string `json:"distro_description"`
	DistroVersion     string `json:"distro_version"`
	FrontendConfig0   string `json:"frontend_config#0"`
	FrontendType0     string `json:"frontend_type#0"`
	Hostname          string `json:"hostname"`
	ID                string `json:"id"`
	KernelDescription string `json:"kernel_description"`
	KernelVersion     string `json:"kernel_version"`
	MemSwapKb         string `json:"mem_swap_kb"`
	MemTotalKb        string `json:"mem_total_kb"`
	NumHandles        string `json:"num_handles"`
	OS                string `json:"os"`
	PID               string `json:"pid"`
	RealmID           string `json:"realm_id"`
	RealmName         string `json:"realm_name"`
	ZoneID            string `json:"zone_id"`
	ZoneName          string `json:"zone_name"`
	ZonegroupID       string `json:"zonegroup_id"`
	ZonegroupName     string `json:"zonegroup_name"`
}

type ReportServiceMapServicesRgwDaemon struct {
	StartEpoch int                                       `json:"start_epoch"`
	StartStamp string                                    `json:"start_stamp"`
	Gid        int                                       `json:"gid"`
	Addr       string                                    `json:"addr"`
	Metadata   ReportServiceMapServicesRgwDaemonMetadata `json:"metadata"`
	TaskStatus struct{}                                  `json:"task_status"`
}

type ReportServiceMapServicesRgwDaemonGenericMap map[string]ReportServiceMapServicesRgwDaemon

func (s *ReportServiceMapServicesRgwDaemonGenericMap) UnmarshalJSON(in []byte) error {
	intermediate := map[string]any{}
	if err := json.Unmarshal(in, &intermediate); err != nil {
		return err
	}

	res := map[string]ReportServiceMapServicesRgwDaemon{}
	for k, v := range intermediate {
		if k == "summary" {
			continue
		}

		j, err := json.Marshal(v)
		if err != nil {
			return err
		}

		out := ReportServiceMapServicesRgwDaemon{}
		if err := json.Unmarshal(j, &out); err != nil {
			return err
		}

		res[k] = out
	}

	*s = res

	return nil
}

type ReportServiceMapServicesRgw struct {
	Daemons ReportServiceMapServicesRgwDaemonGenericMap `json:"daemons"`
}

type ReportServiceMapServices struct {
	OSD ReportServiceMapServicesOSD `json:"osd"`
	Rgw ReportServiceMapServicesRgw `json:"rgw"`
}

type ReportServiceMap struct {
	Epoch    int                      `json:"epoch"`
	Modified string                   `json:"modified"`
	Services ReportServiceMapServices `json:"services"`
}

type Report struct {
	ClusterFingerprint    string                       `json:"cluster_fingerprint"`
	Version               string                       `json:"version"`
	Commit                string                       `json:"commit"`
	Timestamp             string                       `json:"timestamp"`
	Tag                   string                       `json:"tag"`
	Health                ReportHealth                 `json:"health"`
	MonmapFirstCommitted  int                          `json:"monmap_first_committed"`
	MonmapLastCommitted   int                          `json:"monmap_last_committed"`
	MonMap                ReportMonMap                 `json:"monmap"`
	Quorum                []int                        `json:"quorum"`
	OSDMap                ReportOSDMap                 `json:"osdmap"`
	OSDMetadata           []ReportOSDMetadata          `json:"osd_metadata"`
	OSDMapCleanEpochs     ReportOSDMapCleanEpochs      `json:"osdmap_clean_epochs"`
	OSDMapFirstCommitted  int                          `json:"osdmap_first_committed"`
	OSDMapLastCommitted   int                          `json:"osdmap_last_committed"`
	CRUSHMap              ReportCRUSHMap               `json:"crushmap"`
	FSMap                 ReportFSMap                  `json:"fsmap"`
	MDSmapFirstCommitted  int                          `json:"mdsmap_first_committed"`
	MDSmapLastCommitted   int                          `json:"mdsmap_last_committed"`
	Auth                  ReportAuth                   `json:"auth"`
	NumPG                 uint32                       `json:"num_pg"`
	NumPGActive           uint32                       `json:"num_pg_active"`
	NumPGUnknown          uint32                       `json:"num_pg_unknown"`
	NumOSD                uint16                       `json:"num_osd"`
	PoolSum               ReportPoolSum                `json:"pool_sum"`
	OSDSum                ReportOSDSum                 `json:"osd_sum"`
	OSDSumByClass         map[string]ReportOSDSumClass `json:"osd_sum_by_class"`
	PoolStats             []ReportPoolStats            `json:"pool_stats"`
	OSDStats              []ReportOSDStats             `json:"osd_stats"`
	NumPGByState          []ReportNumPGByState         `json:"num_pg_by_state"`
	NumPGByOSD            []ReportNumPGByOSD           `json:"num_pg_by_osd"`
	PurgedSnaps           []ReportPurgedSnaps          `json:"purged_snaps"`
	ServiceMap            ReportServiceMap             `json:"servicemap"`
	MgrstatFirstCommitted int                          `json:"mgrstat_first_committed"`
	MgrstatLastCommitted  int                          `json:"mgrstat_last_committed"`
	LogmFirstCommitted    int                          `json:"logm_first_committed"`
	LogmLastCommitted     int                          `json:"logm_last_committed"`
	Paxos                 ReportPaxos                  `json:"paxos"`
}

func (r *Report) ToSvc() (models.ClusterReport, error) {
	crh, err := NewClusterStatusHealth(r.Health.Status)
	if err != nil {
		return models.ClusterReport{}, err
	}

	checks, err := NewClusterStatusCheck(r.Health.Checks)
	if err != nil {
		return models.ClusterReport{}, err
	}

	mutes, err := NewClusterMutedChecks(r.Health.Mutes)
	if err != nil {
		return models.ClusterReport{}, err
	}

	numMons := len(r.MonMap.Mons)
	if numMons > 2^8-1 {
		return models.ClusterReport{}, errors.Wrapf(ErrOverflow, "too many monitors found: %d", numMons)
	}

	numMonsInQuorum := len(r.Quorum)
	if numMonsInQuorum > 2^8-1 {
		return models.ClusterReport{}, errors.Wrapf(ErrOverflow, "too many monitors in quorum found: %d", numMons)
	}

	numPools := len(r.OSDMap.Pools)
	if numPools > 2^16-1 {
		return models.ClusterReport{}, errors.Wrapf(ErrOverflow, "too many pools found: %d", numPools)
	}

	total, osdsUp, osdsIn, osdsWithoutClusterAddress := countOSDs(r.OSDMap.OSDs)

	numPGs, numPGsByState, err := countPGs(r.NumPGByState)
	if err != nil {
		return models.ClusterReport{}, err
	}

	osdDaemons := []models.OSDDaemon{}
	for _, osd := range r.OSDMetadata {
		frontIP, err := parseCephIPAddress(osd.FrontAddr)
		if err != nil {
			return models.ClusterReport{}, errors.Wrap(err, "error parsing front_addr")
		}

		backIP, err := parseCephIPAddress(osd.BackAddr)
		if err != nil {
			return models.ClusterReport{}, errors.Wrap(err, "error parsing back_addr")
		}

		memoryTotalBytes, err := strconv.ParseUint(osd.MemTotalKb, 10, 64)
		if err != nil {
			return models.ClusterReport{}, errors.Wrap(err, "error parsing mem_total_kb value")
		}

		swapTotalBytes, err := strconv.ParseUint(osd.MemSwapKb, 10, 64)
		if err != nil {
			return models.ClusterReport{}, errors.Wrap(err, "error parsing mem_swap_kb value")
		}

		isRotational, err := strconv.ParseBool(osd.Rotational)
		if err != nil {
			return models.ClusterReport{}, errors.Wrap(err, "error parsing rotational value")
		}

		osdDaemons = append(osdDaemons, models.OSDDaemon{
			ID:               uint16(osd.ID),
			Hostname:         osd.Hostname,
			Architecture:     osd.Arch,
			FrontIP:          frontIP,
			BackIP:           backIP,
			MemoryTotalBytes: memoryTotalBytes,
			SwapTotalBytes:   swapTotalBytes,
			IsRotational:     isRotational,
			Devices:          strings.Split(osd.Devices, ","),
		})
	}

	return models.ClusterReport{
		HealthStatus:                 crh,
		Checks:                       checks,
		MutedChecks:                  mutes,
		NumMons:                      uint8(numMons),
		NumMonsInQuorum:              uint8(numMonsInQuorum),
		AllowCrimson:                 r.OSDMap.AllowCrimson,
		StretchMode:                  r.MonMap.StretchMode,
		NumOSDs:                      total,
		NumOSDsWithoutClusterAddress: osdsWithoutClusterAddress,
		NumOSDsIn:                    osdsIn,
		NumOSDsUp:                    osdsUp,
		NumOSDsByRelease:             countOSDsByRelease(r.OSDMetadata),
		NumOSDsByVersion:             countOSDsByVersion(r.OSDMetadata),
		NumOSDsByDeviceType:          countOSDsByDeviceType(r.OSDMetadata),
		OSDDaemons:                   osdDaemons,
		TotalOSDCapacityKB:           r.OSDSum.Kb,
		TotalOSDUsedDataKB:           r.OSDSum.KbUsedData,
		TotalOSDUsedMetaKB:           r.OSDSum.KbUsedMeta,
		TotalOSDUsedOMAPKB:           r.OSDSum.KbUsedOmap,
		NumPools:                     uint16(numPools),
		NumPGs:                       numPGs,
		NumPGsByState:                numPGsByState,
		BackfillfullRatio:            r.OSDMap.BackfillfullRatio,
		FullRatio:                    r.OSDMap.FullRatio,
		NearfullRatio:                r.OSDMap.NearfullRatio,
		RequireMinCompatClient:       r.OSDMap.MinCompatClient,
	}, nil
}

func parseCephIPAddress(in string) (string, error) {
	addr := strings.SplitN(in, ":", 2)
	if len(addr) != 2 {
		return "", errors.New("malformed ceph address string")
	}

	addreses := strings.Split(addr[1], ",")
	if len(addreses) != 2 {
		return "", errors.New("malformed ceph address representation")
	}

	addressParts := strings.SplitN(addreses[0], ":", 2)
	if len(addressParts) != 2 {
		return "", errors.New("malformed ceph address string")
	}

	return addressParts[0], nil
}

func countOSDs(osds []ReportOSDMapOSD) (total, up, in, withoutClusterAddress uint16) {
	for _, osd := range osds {
		total++

		if osd.In == 1 {
			in++
		}

		if osd.Up == 1 {
			up++
		}

		if len(osd.ClusterAddrs.Addrvec) == 0 {
			withoutClusterAddress++
		}
	}
	return
}

func countOSDsByRelease(osds []ReportOSDMetadata) map[string]uint16 {
	c := make(map[string]uint16)
	for _, r := range osds {
		if _, ok := c[r.CephRelease]; !ok {
			c[r.CephRelease] = 0
		}

		c[r.CephRelease]++
	}
	return c
}

func countOSDsByVersion(osds []ReportOSDMetadata) map[string]uint16 {
	c := make(map[string]uint16)
	for _, r := range osds {
		if _, ok := c[r.CephVersionShort]; !ok {
			c[r.CephVersionShort] = 0
		}

		c[r.CephVersionShort]++
	}
	return c
}

func countOSDsByDeviceType(osds []ReportOSDMetadata) map[string]uint16 {
	c := make(map[string]uint16)
	for _, r := range osds {
		if _, ok := c[r.BluestoreBdevType]; !ok {
			c[r.BluestoreBdevType] = 0
		}

		c[r.BluestoreBdevType]++
	}
	return c
}

func countPGs(pgs []ReportNumPGByState) (total uint32, byState map[string]uint32, err error) {
	pgsByState := make(map[string]uint32)
	for _, pg := range pgs {
		states := strings.Split(pg.State, "+")
		for _, state := range states {
			if _, ok := pgsByState[state]; !ok {
				pgsByState[state] = 0
			}
			pgsByState[state] += pg.Num
		}
		total += pg.Num
	}

	return total, pgsByState, nil
}

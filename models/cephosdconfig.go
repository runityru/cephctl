package models

type CephOSDConfig struct {
	AllowCrimson           bool    `yaml:"allow_crimson" diff:"allow_crimson" default:"false"`
	BackfillfullRatio      float32 `yaml:"backfillfull_ratio" diff:"backfillfull_ratio" default:"0.9"`
	FullRatio              float32 `yaml:"full_ratio" diff:"full_ratio" default:"0.95"`
	NearfullRatio          float32 `yaml:"nearfull_ratio" diff:"nearfull_ratio" default:"0.85"`
	RequireMinCompatClient string  `yaml:"require_min_compat_client" diff:"require_min_compat_client" default:"reef"`
}

type CephOSDConfigDifferenceKind string

type CephOSDConfigDifference struct {
	Key      string
	OldValue string
	Value    string
}

package models

type ConfigOption struct {
	Section            string `json:"section"`
	Name               string `json:"name"`
	Value              string `json:"value"`
	Level              string `json:"level"`
	CanUpdateAtRuntime bool   `json:"can_update_at_runtime"`
	Mask               string `json:"mask"`
}

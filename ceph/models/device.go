package models

type DeviceLocation struct {
	Host string `json:"host"`
	Dev  string `json:"dev"`
	Path string `json:"path"`
}

type Device struct {
	DevID     string           `json:"devid"`
	Location  []DeviceLocation `json:"location"`
	Daemons   []string         `json:"daemons"`
	WearLevel float64          `json:"wear_level"`
}

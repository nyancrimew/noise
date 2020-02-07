package wetter_com

type Record struct {
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	VerticalAccuracy   string  `json:"vertical_accuracy"`
	UtcTimestamp       int64   `json:"utc_timestamp"`
	Speed              float64 `json:"speed"`
	SpeedAccuracy      string  `json:"speed_accuracy"`
	Course             float64 `json:"course"`
	CourseAccuracy     string  `json:"course_accuracy"`
	Altitude           int     `json:"altitude"`
	LocationSource     string  `json:"location_source"`
	AppVersionCode     int64   `json:"app_version_code"`
	DeviceManufacturer string  `json:"device_manufacturer"`
	DeviceModel        string  `json:"device_model"`
	OsVersion          string  `json:"os_version"`
	QueryAccuracy      string  `json:"query_accuracy"`
	CarrierName        string  `json:"carrier_name"`
	WifiSsid           string  `json:"wifi_ssid"`
	WifiBssid          string  `json:"wifi_bssid"`
	ConfigHash         string  `json:"config_hash"`
	LocationMethod     string  `json:"location_method"`
	LocationContext    string  `json:"location_context"`
	AdID               string  `json:"ad_id"`
}
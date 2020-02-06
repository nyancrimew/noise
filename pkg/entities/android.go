package entities

import "github.com/deletescape/noise/internal/data/android"

type AndroidDevice struct {
	Manufacturer string `json:"manufacturer"`
	MarketName   string `json:"market_name"`
	Codename     string `json:"codename"`
	Model        string `json:"model"`
}

type AndroidVersion struct {
	Release android.Release
	Code    android.VersionCode
}

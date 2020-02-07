package data

import (
	"github.com/deletescape/noise/internal/data/android"
	"github.com/deletescape/noise/pkg/entities"
	"github.com/deletescape/noise/pkg/wyrand"
)

func GetAndroidDevice() entities.AndroidDevice {
	return android.Devices[wyrand.IntN(len(android.Devices))]
}

func GetAndroidVersion(min entities.AndroidVersionCode, max entities.AndroidVersionCode) entities.AndroidVersion {
	version := entities.AndroidVersionCode(wyrand.IntN(int(max-min)) + int(min))
	releases := version.GetReleases()
	return entities.AndroidVersion{
		Release: releases[wyrand.IntN(len(releases))],
		Code:    version,
	}
}
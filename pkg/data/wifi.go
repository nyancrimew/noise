package data

import (
	"fmt"
	"github.com/deletescape/noise/internal/data/ieee"
	"github.com/deletescape/noise/internal/data/wifi"
	"github.com/deletescape/noise/pkg/entities"
	"github.com/deletescape/noise/pkg/wyrand"
)

func GetBssid() string {
	oui := ieee.OUIs[wyrand.IntN(len(ieee.OUIs))]
	return fmt.Sprintf("%s%x:%x:%x", oui, wyrand.Uint8(), wyrand.Uint8(), wyrand.Uint8())
}

func GetWiFi() entities.WiFi {
	return entities.WiFi{
		SSID:  wifi.SSIDS[wyrand.IntN(len(wifi.SSIDS))],
		BSSID: GetBssid(),
	}
}

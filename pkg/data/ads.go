package data

import "github.com/deletescape/noise/pkg/wyrand"

func GetAdId() string {
	return wyrand.UUID3().String()
}

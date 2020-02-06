package data

import (
	"github.com/deletescape/noise/internal/data/carriers"
	"github.com/deletescape/noise/pkg/entities"
	"github.com/deletescape/noise/pkg/wyrand"
)

func GetApn() entities.Apn {
	return carriers.ApnConfigs[wyrand.IntN(len(carriers.ApnConfigs))]
}

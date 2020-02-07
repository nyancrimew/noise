package data

import (
	"github.com/deletescape/noise/pkg/entities"
	"github.com/deletescape/noise/pkg/wyrand"
)

// TODO: ensure that altitude actually correlates with location?
func GetLocation() entities.Location {
	var location entities.Location
	// 0 = north america, 1 = south america, 2 = africa, 3 = europe, 4 = asia, 5 oceania
	continent := wyrand.Uint8N(6)
	switch continent {
	case 0:
		location.Lat = -(wyrand.Float64N(82) + 51)
		location.Lon = wyrand.Float64N(54) + 13
		break
	case 1:
		location.Lat = -(wyrand.Float64N(41) + 37)
		location.Lon = wyrand.Float64N(72) - 50
		break
	case 2:
		location.Lat = wyrand.Float64N(60) - 8
		location.Lon = wyrand.Float64N(63) - 35
		break
	case 3:
		location.Lat = wyrand.Float64N(76) - 10
		location.Lon = wyrand.Float64N(40) + 32
		break
	case 4:
		location.Lat = wyrand.Float64N(76) - 10
		location.Lon = wyrand.Float64N(40) + 32
		break
	case 5:
		location.Lat = wyrand.Float64N(127) + 34
		location.Lon = wyrand.Float64N(58) + 12
		break
	}

	location.Altitude = wyrand.Float64N(2200) - 200

	return location
}

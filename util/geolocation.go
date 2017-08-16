package util

import (
	"math"
)

var (
	deg2rad       = math.Pi / 180 // degrees to radian conversion
	rad2deg       = 180 / math.Pi // radians to degrees conversion
	earthRadius   = 6371.01       // Earth's radius in km
	maxLat        = math.Pi / 2   // 90 degrees
	minLat        = -maxLat       // -90 degrees
	maxLon        = math.Pi       // 180 degrees
	minLon        = -maxLon       // -180 degrees
	fullCircleRad = math.Pi * 2   // Full cirle (360 degrees) in radians
)

func degtorad(value float64) float64 {
	return value * deg2rad
}

func radtodeg(value float64) float64 {
	return value * rad2deg
}

func fromDegrees(lat float64, lon float64) (radLat float64, radLon float64) {
	return degtorad(lat), degtorad(lon)
}

func fromRadians(lat float64, lon float64) (radLat float64, radLon float64) {
	return radtodeg(lat), radtodeg(lon)
}

func distanceTo(lat float64, lon float64, pointLat float64, pointLon float64) float64 {
	return math.Acos(math.Sin(pointLat)*math.Sin(lat)+
		math.Cos(pointLat)*math.Cos(lat)*
			math.Cos(pointLon-lon)) * earthRadius
}

// BoundingCoordinates : calculate the bounding coordinates for a specific point and distance
func BoundingCoordinates(lat float64, lon float64, distance float64) [4]float64 {
	var radDist = distance / earthRadius // angular distance in radians on a great circle
	var radLat = degtorad(lat)
	var radLon = degtorad(lon)
	var relMinLat = radLat - radDist
	var relMaxLat = radLat + radDist

	var relMinLon, relMaxLon float64

	if relMinLat > minLat && relMaxLat < maxLat {
		var deltaLon = math.Asin(math.Sin(radDist) / math.Cos(radLat))
		relMinLon = radLon - deltaLon

		if relMinLon < minLon {
			relMinLon += 2 * math.Pi
		}

		relMaxLon = radLon + deltaLon
		if relMaxLon > maxLon {
			relMaxLon -= 2 * math.Pi
		}
	} else {
		// a pole is within the distance
		relMinLat = math.Max(relMinLat, minLat)
		relMaxLat = math.Min(relMaxLat, maxLat)
		relMinLon = minLon
		relMaxLon = maxLon
	}

	radMinLat, radMinLon := fromRadians(relMinLat, relMinLon)
	radMaxLat, radMaxLon := fromRadians(relMaxLat, relMaxLon)

	return [4]float64{radMinLat, radMinLon, radMaxLat, radMaxLon}
}

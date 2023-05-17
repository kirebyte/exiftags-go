package entities

import "fmt"

type GpsInfo struct {
	Latitude, Longitude GpsLocation
}

type GpsLocation struct {
	Orientation byte
	Degrees     int
	Minutes     int
	Seconds     float64
}

func (g *GpsLocation) String() string {
	return fmt.Sprintf("%d°%d'%f\"%s", g.Degrees, g.Minutes, g.Seconds, string(g.Orientation))
}

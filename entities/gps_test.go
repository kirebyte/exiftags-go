package entities

import (
	"testing"
)

func TestGpsLocationString(t *testing.T) {
	cases := []struct {
		in   GpsLocation
		want string
	}{
		{GpsLocation{'N', 40, 26, 46.52}, "40°26'46.520000\"N"},
		{GpsLocation{'S', 23, 30, 31.00}, "23°30'31.000000\"S"},
		{GpsLocation{'E', 100, 0, 0.00}, "100°0'0.000000\"E"},
		{GpsLocation{'W', 77, 3, 23.00}, "77°3'23.000000\"W"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("GpsLocation.String() == %q, want %q", got, c.want)
		}
	}
}

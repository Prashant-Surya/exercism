package space

import (
	"math"
)

type Planet string

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func Age(seconds float64, planet Planet) float64 {
	var days float64 = float64(seconds) / 86400
	var earth_years float64 = days / 365.25
	dmap := map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	if string(planet) == "Earth" {
		return Round(earth_years, .5, 2)
	}

	return Round(earth_years/dmap[string(planet)], .5, 2)
}

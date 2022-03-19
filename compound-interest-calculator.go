package main

import (
	"math"
)

func calCi(capital float64, rate float64, year int) float64 {
	return capital * math.Pow(1.0+rate, float64(year))
}

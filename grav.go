package main

import "math"

func gravForce(m1 float64, m2 float64, r float64) float64 {

	// Newton's constant
	G := 6.674 * (math.Pow(10, -11))

	formula := G * (m1 * m2 / r)

	return math.Round(formula)

}

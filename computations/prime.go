package computations

import "math"

// IsPrime - func
func IsPrime(value int) interface{} {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return "false"
		}
	}
	if value > 1 {
		return "true"
	}
	return "false"
}

package computations

import (
	"math"
)

// IsPrime - func
func IsPrime(n int) string {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return "false"
		}
	}
	if n > 1 {
		return "true"
	}
	return "false"
}

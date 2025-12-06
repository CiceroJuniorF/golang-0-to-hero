package tasks

import (
	"fmt"
	"math"
)

func ConditionalSqrt(x float64) string {
	if x < 0 {
		return ConditionalSqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ShortDeclarationPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func ShortDeclarationPowWithElseAlert(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("Short Declaration Pow With Else %g >= %g \n", v, lim)
	}
	return lim
}

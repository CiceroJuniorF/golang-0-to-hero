package tasks

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	Number float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e.Number)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &ErrNegativeSqrt{x}
	}
	return 0, nil
}

func RunError() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

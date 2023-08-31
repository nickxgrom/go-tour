package Exercises

import (
	"math"
)

type ErrNegativeSqrt struct{}

func (e ErrNegativeSqrt) Error() string {
	return "the given value is negative"
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt{}
	}

	return math.Sqrt(x), nil
}

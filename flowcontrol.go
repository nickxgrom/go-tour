package main

import (
	"fmt"
	"math"
)

func Sqrt(x, precision float64) float64 {
	i := 0
	z := 1.0
	previousVal := precision
	for math.Abs(z-previousVal) >= precision {
		val := (z*z - x) / (2 * z)
		previousVal = z
		z -= val
		i++
	}

	return z
}

func DeferTest() {
	num := 1

	defer fmt.Println(num)

	num++

	fmt.Println(num)
}

func DeferStack() {
	i := 0

	for i < 10 {
		defer fmt.Println(i)
		i++
	}
}

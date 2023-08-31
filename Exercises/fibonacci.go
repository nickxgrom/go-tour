package Exercises

// Fibonacci TODO: try with prev1, prev2 values
func Fibonacci() func() int {
	prevVal := 0
	currentVal := 1
	nextVal := 0

	return func() int {
		defer func() {
			nextVal = prevVal + currentVal
			prevVal = currentVal
			currentVal = nextVal
		}()

		return prevVal
	}
}

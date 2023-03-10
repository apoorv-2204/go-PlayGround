package naive_math

// Babylonian method. It starts with an initial guess and then updates it repeatedly until it reaches a desired level of accuracy.
func squareRoot(number float64, precision int) float64 {
	var result float64 = 1.0
	for i := 0; i < precision; i++ {

		result = 0.5 * (result + number/result)
	}
	return result
}

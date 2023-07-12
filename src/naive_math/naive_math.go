package naive_math

var logMessage = "[LOG][naive_math]"

// The logMessage variable can be called only from within the package

// Version of the naive_math
// Version variable can be reached from anywhere
var Version string = "0.1.0"

// Babylonian method. It starts with an initial guess and then updates it repeatedly until it reaches a desired level of accuracy.
func SquareRoot(number float64, precision int) float64 {
	var result float64 = 1.0
	for i := 0; i < precision; i++ {

		result = 0.5 * (result + number/result)
	}
	return result
}

// go build

package exercism

var Version string = "v0.1.0"

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return 2 * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return 2*numberOfLayers + actualMinutesInOven
}

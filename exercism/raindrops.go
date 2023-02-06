package exercism

import (
	"fmt"
)

const sound_of_3 = "Pling"
const sound_of_5 = "Plang"
const sound_of_7 = "Plong"

func Convert(number int) string {
	var result string = ""
	if number%3 == 0 {
		result += sound_of_3
	}
	if number%5 == 0 {
		result += sound_of_5
	}
	if number%7 == 0 {
		result += sound_of_7
	}
	if result == "" {
		result += fmt.Sprintf("%d", number)
	}

	return result

}

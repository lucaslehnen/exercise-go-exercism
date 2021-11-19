package raindrops

import (
	"strconv"
)

// Convert a number into a sound
func Convert(number int) string {
	factors := [3]int{3, 5, 7}
	sounds := [3]string{"Pling", "Plang", "Plong"}

	var resultSound string

	for index, factor := range factors {
		if number%factor == 0 {
			resultSound += sounds[index]
		}
	}
	if resultSound == "" {
		resultSound = strconv.Itoa(number)
	}
	return resultSound
}

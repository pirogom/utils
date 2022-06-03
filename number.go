package utils

import "strconv"

/**
*	atoi
**/
func Atoi(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		return 0
	}

	return num
}

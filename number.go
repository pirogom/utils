package utils

import "strconv"

/**
*	atoi
**/
func atoi(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		return 0
	}

	return num
}

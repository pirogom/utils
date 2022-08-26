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

/**
*	atoi64
**/
func Atoi64(sv string) int64 {
	v, e := strconv.ParseInt(sv, 10, 64)
	if e != nil {
		return 0
	}
	return v
}

/**
*	atof64
**/
func Atof64(sv string) float64 {
	f, ferr := strconv.ParseFloat(sv, 64)

	if ferr != nil {
		return 0.0
	}
	return f
}

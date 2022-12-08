package helpers

import "strconv"

// ParseInt is used to parse a string to an integer. ParseInt panics if the given string cannot be parse to an integer.
func ParseInt(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return val
}

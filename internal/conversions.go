package internal

import (
	"strconv"
)

func StringToIntFast(input string) int {
	i, _ := strconv.Atoi(input)
	return i
}

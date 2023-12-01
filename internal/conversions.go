package internal

import (
	"fmt"
	"strconv"
)

func StringToIntFast(input string) int {
	i, _ := strconv.Atoi(input)
	return i
}

func FastPrint(input any) {
	fmt.Println(input)
}

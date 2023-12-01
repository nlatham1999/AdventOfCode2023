package days2018

import (
	"fmt"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunForDayOne() {
	inputs, _ := internal.ReadFileLineByLine("./inputs2018/day1")
	total := 0
	resultingFreq := make(map[int]interface{})
	resultingFreq[total] = nil
	for true {
		fmt.Println("round")
		for _, x := range inputs {
			// fmt.Println(x)
			first := x[:1]
			second := x[1:]
			if first == "+" {
				total += internal.StringToIntFast(second)
			} else {
				total -= internal.StringToIntFast(second)
			}
			if _, found := resultingFreq[total]; found {
				fmt.Println(total)
				return
			}
			resultingFreq[total] = nil
		}
	}

	fmt.Println(resultingFreq)

	fmt.Println(total)
}

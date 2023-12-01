package days2022

import (
	"fmt"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayTwo() {
	inputs, err := internal.ReadFileLineByLine("./inputs2022/day2")
	if err != nil {
		fmt.Println("Could not open")
	}

	for _, x := range inputs {
		r := []rune(x)
		fmt.Println(r)
		//etc etc
	}
}

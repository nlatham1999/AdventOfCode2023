package days2022

import (
	"fmt"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayOne() {

	lines, err := internal.ReadFileLineByLine("./inputs/day1")
	if err != nil {
		fmt.Println("Got an error reading the file")
	}

	localMax := 0
	amounts := []int{}
	for _, x := range lines {
		if x == "" {
			amounts = append(amounts, localMax)
			localMax = 0
		} else {
			localMax += internal.StringToIntFast(x)
		}
	}
	internal.SortIntDesc(amounts)

	fmt.Println(internal.SumOfSliceInt(amounts[:3]))

	fmt.Println(amounts[:3])
}

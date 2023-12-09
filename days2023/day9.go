package days2023

import (
	"fmt"
	"strings"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayNine() {
	DayNinePartOne()
	DayNinePartTwo()
}

func DayNinePartTwo() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day9.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	total := 0
	for _, row := range inputs {
		r := strings.Split(row, " ")

		values := [][]int{}
		v := []int{}
		for _, x := range r {
			v = append(v, internal.StringToIntFast(x))
		}
		values = append(values, v)

		//construct the sequences
		allZeros := false
		for !allZeros {
			previousRow := values[len(values)-1]
			newRow := []int{}
			for i := 0; i < len(previousRow); i++ {
				if i > 0 {
					newRow = append(newRow, previousRow[i]-previousRow[i-1])
				}
			}

			//check for zeros
			allZeros = true
			for _, x := range newRow {
				if x != 0 {
					allZeros = false
				}
			}

			values = append(values, newRow)
		}

		//get the value
		sum := 0
		for i := len(values) - 1; i >= 0; i-- {
			if i > 0 {
				aboveRow := values[i-1]
				sum = aboveRow[0] - sum
			}
		}

		total += sum

		// fmt.Println(total)
	}

	fmt.Println(total)
}

func DayNinePartOne() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day9.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	total := 0
	for _, row := range inputs {
		r := strings.Split(row, " ")

		values := [][]int{}
		v := []int{}
		for _, x := range r {
			v = append(v, internal.StringToIntFast(x))
		}
		values = append(values, v)

		//construct the sequences
		allZeros := false
		for !allZeros {
			previousRow := values[len(values)-1]
			newRow := []int{}
			for i := 0; i < len(previousRow); i++ {
				if i > 0 {
					newRow = append(newRow, previousRow[i]-previousRow[i-1])
				}
			}

			//check for zeros
			allZeros = true
			for _, x := range newRow {
				if x != 0 {
					allZeros = false
				}
			}

			values = append(values, newRow)
		}

		//get the value
		sum := 0
		for i := len(values) - 1; i >= 0; i-- {
			if i > 0 {
				aboveRow := values[i-1]
				sum = sum + aboveRow[len(aboveRow)-1]
			}
		}

		total += sum

		// fmt.Println(total)
	}

	fmt.Println(total)
}

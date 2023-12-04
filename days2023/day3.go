package days2023

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayThree() {

	ThreePartOne()
	ThreePartTwo()
}

func ThreePartTwo() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day3.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	fmt.Println(len([]rune(inputs[0])))
	fmt.Println(len(inputs))

	totalSum := 0

	foundPieces := map[string][]int{}

	for i, row := range inputs {
		inNum := false
		numStr := ""
		symbolFound := false
		symbolPositions := make(map[string]interface{})
		for j, c := range row {
			if unicode.IsDigit(c) {
				if !inNum {
					inNum = true
				}
				numStr += string(c)

				//check if next to a symbol
				x_min := i - 1
				if x_min < 0 {
					x_min = 0
				}
				x_max := i + 1
				if x_max == len(inputs) {
					x_max -= 1
				}
				y_min := j - 1
				if y_min < 0 {
					y_min = 0
				}
				y_max := j + 1
				if y_max == len([]rune(inputs[j])) {
					y_max -= 1
				}

				for ii := x_min; ii < x_max+1; ii++ {
					for iii := y_min; iii < y_max+1; iii++ {
						r := []rune(inputs[ii])[iii]
						if r == '*' {
							symbolFound = true
							symbolPositions[strconv.Itoa(ii)+"_"+strconv.Itoa(iii)] = nil
						}
					}
				}

			} else {
				inNum = false
				if symbolFound {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println("Could not parse number. " + numStr)
						return
					}
					for p, _ := range symbolPositions {
						_, found := foundPieces[p]
						if found {
							foundPieces[p] = append(foundPieces[p], num)
						} else {
							foundPieces[p] = []int{num}
						}
					}

				}
				for k := range symbolPositions {
					delete(symbolPositions, k)
				}

				symbolFound = false
				numStr = ""
			}
		}

		if symbolFound {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Could not parse number. " + numStr)
				return
			}
			for p, _ := range symbolPositions {
				_, found := foundPieces[p]
				if found {
					foundPieces[p] = append(foundPieces[p], num)
				} else {
					foundPieces[p] = []int{num}
				}
			}
		}
	}

	for _, val := range foundPieces {
		if len(val) == 2 {
			totalSum += val[0] * val[1]
		}
	}

	fmt.Println(totalSum)
}

func ThreePartOne() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day3.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	fmt.Println(len([]rune(inputs[0])))
	fmt.Println(len(inputs))

	totalSum := 0
	for i, row := range inputs {
		inNum := false
		numStr := ""
		symbolFound := false
		for j, c := range row {
			if unicode.IsDigit(c) {
				if !inNum {
					inNum = true
				}
				numStr += string(c)

				//check if next to a symbol
				x_min := i - 1
				if x_min < 0 {
					x_min = 0
				}
				x_max := i + 1
				if x_max == len(inputs) {
					x_max -= 1
				}
				y_min := j - 1
				if y_min < 0 {
					y_min = 0
				}
				y_max := j + 1
				if y_max == len([]rune(inputs[j])) {
					y_max -= 1
				}

				for ii := x_min; ii < x_max+1; ii++ {
					for iii := y_min; iii < y_max+1; iii++ {
						r := []rune(inputs[ii])[iii]
						if r != '.' && !unicode.IsDigit(r) {
							symbolFound = true
						}
					}
				}

			} else {
				inNum = false
				if symbolFound {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println("Could not parse number. " + numStr)
						return
					}
					totalSum += num
				}
				symbolFound = false
				numStr = ""
			}
		}

		if symbolFound {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Could not parse number. " + numStr)
				return
			}
			totalSum += num
		}
	}

	fmt.Println(totalSum)
}

package days2023

import (
	"fmt"
	"unicode"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayOne() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day1")
	if err != nil {
		fmt.Println("Could not open")
	}

	sum := 0
	for _, x := range inputs {
		r := []rune(x)

		first := -1
		second := -1
		for i, y := range r {

			digit := -1
			if unicode.IsDigit(y) {
				digit = int(y - '0')

			}

			if len(r) > i+2 {
				if stringToNum(string(r[i:i+3])) != -1 {
					digit = stringToNum(string(r[i : i+3]))
				}

			}

			if len(r) > i+3 {
				if stringToNum(string(r[i:i+4])) != -1 {
					digit = stringToNum(string(r[i : i+4]))
				}
			}

			if len(r) > i+4 {
				if stringToNum(string(r[i:i+5])) != -1 {
					digit = stringToNum(string(r[i : i+5]))
				}
			}

			if digit != -1 {
				if first == -1 {
					first = digit
				} else {
					second = digit
				}
			}
		}
		if second == -1 {
			second = first
		}

		sum += first*10 + second

		fmt.Println(first*10 + second)

		//etc etc
	}

	fmt.Println(sum)
}

func stringToNum(str string) int {
	switch str {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1
	}
}

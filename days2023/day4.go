package days2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayFour() {

	FourPartOne()
	FourPartTwo()
}

func FourPartTwo() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day4.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	occurences := map[int]int{}

	sum := 0
	for index, row := range inputs {

		_, f := occurences[index]
		if !f {
			occurences[index] = 1
		}

		s := strings.Split(row, ":")
		rounds := strings.Split(s[1], "|")

		//load first set into checker
		checker := make(map[int]interface{})
		nums := strings.Split(rounds[0], " ")
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err == nil {
				checker[i] = nil
			}
		}

		nums2 := strings.Split(rounds[1], " ")
		matchCount := 0
		for _, num := range nums2 {
			i, err := strconv.Atoi(num)
			if err == nil {
				if _, found := checker[i]; found {
					matchCount += 1
				}
			}
		}

		for i := 0; i < matchCount; i++ {
			j := index + i + 1
			if j < len(inputs) {
				roundOcc, _ := occurences[j]
				if roundOcc == 0 {
					roundOcc = 1
				}
				occurences[j] = roundOcc + occurences[index]
				// fmt.Println(fmt.Sprintf("From index %d, adding to index %d, now %d", index, j, roundOcc+occurences[index]))
			}
		}

	}

	for _, val := range occurences {
		// fmt.Println(fmt.Sprintf("%d : %d", key, val))
		sum += val
	}

	fmt.Println(sum)
}

func FourPartOne() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day4.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	sum := 0
	for _, row := range inputs {
		score := 0

		s := strings.Split(row, ":")
		rounds := strings.Split(s[1], "|")

		//load first set into checker
		checker := make(map[int]interface{})
		nums := strings.Split(rounds[0], " ")
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err == nil {
				checker[i] = nil
			}
		}

		nums2 := strings.Split(rounds[1], " ")
		for _, num := range nums2 {
			i, err := strconv.Atoi(num)
			if err == nil {
				if _, found := checker[i]; found {
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			}
		}
		sum += score

	}

	fmt.Println(sum)
}

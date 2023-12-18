package days2023

import (
	"fmt"
	"strings"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDay12() {
	// RunDay12PartOne()
	RunDay12PartTwo()
}

var possibilities [][]rune

func RunDay12PartTwo() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day12.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	sum := 0
	for _, row := range inputs {
		t := strings.Split(row, " ")
		r := t[0]

		rules := []int{}
		rulesS := strings.Split(t[1], ",")
		for _, v := range rulesS {
			rules = append(rules, internal.StringToIntFast(v))
		}

		newRules := []int{}
		newRules = append(newRules, rules...)
		newRules = append(newRules, rules...)
		newRules = append(newRules, rules...)
		newRules = append(newRules, rules...)
		newRules = append(newRules, rules...)

		newR := r + "?" + r + "?" + r + "?" + r + "?" + r
		possibilities = [][]rune{}

		constructMissing2([]rune(newR), []rune{}, 0, 0, 0, newRules)

		count := 0
		fmt.Println()
		fmt.Println(len(possibilities))
		for _, p := range possibilities {
			if matchesRules(p, newRules) {
				// fmt.Println(string(p))
				count++
			}
		}

		sum += count
		fmt.Println(newR)
		// fmt.Println(len(possibilities))
		fmt.Println(count)
	}

	fmt.Println(sum)
}

func RunDay12PartOne() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day12.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	sum := 0
	for _, row := range inputs {
		t := strings.Split(row, " ")
		r := t[0]

		rules := []int{}
		rulesS := strings.Split(t[1], ",")
		for _, v := range rulesS {
			rules = append(rules, internal.StringToIntFast(v))
		}
		possibilities = [][]rune{}

		constructMissing([]rune(r), []rune{}, 0)

		count := 0
		fmt.Println(len(possibilities))
		for _, p := range possibilities {
			if matchesRules(p, rules) {
				// fmt.Println(string(p))
				count++
			}
		}

		sum += count
		fmt.Println(r)
		// fmt.Println(len(possibilities))
		fmt.Println(count)
	}

	fmt.Println(sum)
}

func deepCopyRuneSlice(original []rune) []rune {
	copySlice := make([]rune, len(original))
	for i, v := range original {
		copySlice[i] = v
	}
	return copySlice
}

func constructMissing2(row []rune, newRow []rune, x int, currentLength int, currentRule int, rules []int) {

	if currentRule >= len(rules) {
		return
	}

	if !matchesRulesPartially(newRow, row, rules) {
		return
	}

	if x == len(row) {

		// fmt.Println("returning")
		r := deepCopyRuneSlice(newRow)
		possibilities = append(possibilities, r)
		return
	}

	if row[x] != '?' {

		// fmt.Println("adding normal")
		constructMissing2(row, append(newRow, row[x]), x+1, currentLength, currentRule, rules)
		return
	}

	constructMissing2(row, append(newRow, '.'), x+1, currentLength, currentRule, rules)

	constructMissing2(row, append(newRow, '#'), x+1, currentLength, currentRule, rules)

}

func constructMissing(row []rune, newRow []rune, x int) {
	fmt.Println("test2")

	if x == len(row) {
		// fmt.Println("returning")
		r := deepCopyRuneSlice(newRow)
		possibilities = append(possibilities, r)
		return
	}

	if row[x] != '?' {
		// fmt.Println("adding normal")
		constructMissing(row, append(newRow, row[x]), x+1)
		return
	}

	constructMissing(row, append(newRow, '.'), x+1)

	constructMissing(row, append(newRow, '#'), x+1)

}

func matchesRulesPartially(row []rune, oldRow []rune, rules []int) bool {
	currentRule := 0
	currentLength := 0
	for _, v := range row {
		if v == '.' {
			if currentLength != 0 {
				if currentRule >= len(rules) {
					return false
				}
				if currentLength != rules[currentRule] {
					return false
				}
				currentLength = 0
				currentRule++

			}
		} else {
			currentLength++
		}

	}

	if len(oldRow)-len(row) < rules[currentLength]-currentLength {
		return false
	}

	return true
}

func matchesRules(row []rune, rules []int) bool {
	// fmt.Println(rules)
	currentRule := 0
	currentLength := 0
	for _, v := range row {
		if v == '.' {
			if currentLength != 0 {
				if currentRule >= len(rules) {
					// fmt.Println("TEST1 " + string(row))
					return false
				}
				if currentLength != rules[currentRule] {
					// fmt.Println("TEST2 " + string(row))
					return false
				}
				currentLength = 0
				currentRule++

			}
		} else {
			currentLength++
		}

	}
	if currentLength != 0 && currentRule > len(rules)-1 {
		// fmt.Println("TEST3 " + string(row))
		return false
	}
	if currentLength != 0 && currentLength != rules[currentRule] {
		// fmt.Println("TEST4 " + string(row))
		return false
	}
	if currentRule < len(rules)-1 {
		// fmt.Println("TEST5 " + string(row))
		return false
	}
	if currentLength == 0 && currentRule < len(rules) {
		// fmt.Println("TEST6 " + string(row))
		return false
	}

	return true
}

package days2023

import (
	"fmt"
	"strings"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

type Comparer interface {
	greaterThan(a string, b string) bool
}

type ComparerPartOne struct{}

type ComparerPartTwo struct{}

func RunDaySeven() {

	DaySevenPartOne()

	DaySevenPartTwo()

}

func DaySevenPartTwo() {
	c := ComparerPartTwo{}
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day7.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	occurences := map[string]int{}

	keys := []string{}

	for _, line := range inputs {
		p := strings.Split(line, " ")
		occurences[p[0]] = internal.StringToIntFast(p[1])
		keys = append(keys, p[0])
	}

	score := 0
	sorted := mergeSort(keys, c)
	for i, x := range sorted {
		score += occurences[x] * (i + 1)
	}

	fmt.Println(score)
}

func DaySevenPartOne() {
	c := ComparerPartOne{}
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day7.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	occurences := map[string]int{}

	keys := []string{}

	for _, line := range inputs {
		p := strings.Split(line, " ")
		occurences[p[0]] = internal.StringToIntFast(p[1])
		keys = append(keys, p[0])
	}

	score := 0
	sorted := mergeSort(keys, c)
	for i, x := range sorted {
		score += occurences[x] * (i + 1)
	}

	fmt.Println(score)

}

func mergeSort(arr []string, c Comparer) []string {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle], c)
	right := mergeSort(arr[middle:], c)

	return merge(left, right, c)
}

func merge(left, right []string, c Comparer) []string {
	result := []string{}
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if !c.greaterThan(left[leftIndex], right[rightIndex]) {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

// is a greater than b
func (c ComparerPartOne) greaterThan(a string, b string) bool {
	handScoreA := c.handType(a)
	handScoreB := c.handType(b)
	if handScoreA != handScoreB {
		return handScoreA > handScoreB
	}
	rA := []rune(a)
	rB := []rune(b)
	for i := 0; i < 5; i++ {
		nA := c.cardToNum(rA[i])
		nB := c.cardToNum(rB[i])
		if nA != nB {
			return nA > nB
		}
	}
	panic(fmt.Sprintf("Could not compare %s %s", a, b))
}

func (c ComparerPartTwo) greaterThan(a string, b string) bool {
	handScoreA := c.handType(a)
	handScoreB := c.handType(b)
	if handScoreA != handScoreB {
		return handScoreA > handScoreB
	}
	rA := []rune(a)
	rB := []rune(b)
	for i := 0; i < 5; i++ {
		nA := c.cardToNum(rA[i])
		nB := c.cardToNum(rB[i])
		if nA != nB {
			return nA > nB
		}
	}
	panic(fmt.Sprintf("Could not compare %s %s", a, b))
}

func (c *ComparerPartOne) cardToNum(r rune) int {
	if r == 'A' {
		return 14
	}
	if r == 'K' {
		return 13
	}
	if r == 'Q' {
		return 12
	}
	if r == 'J' {
		return 11
	}
	if r == 'T' {
		return 10
	}

	return int(r - '0')
}

func (c ComparerPartTwo) cardToNum(r rune) int {
	if r == 'A' {
		return 14
	}
	if r == 'K' {
		return 13
	}
	if r == 'Q' {
		return 12
	}
	if r == 'J' {
		return 1
	}
	if r == 'T' {
		return 10
	}

	return int(r - '0')
}

func (c ComparerPartOne) handType(hand string) int {
	combined := map[rune]int{}
	for _, r := range hand {
		_, f := combined[r]
		if f {
			combined[r]++
		} else {
			combined[r] = 1
		}
	}

	if len(combined) == 1 {
		return 6 // five of a kind
	}
	if len(combined) == 2 {
		for _, v := range combined {
			if v == 1 || v == 4 {
				return 5 // four of a kind
			} else {
				return 4 // full house
			}
		}
	}
	if len(combined) == 3 {
		twos := 0
		for _, v := range combined {
			if v == 2 {
				twos++
			}
			if v == 3 {
				return 3 // three of a kind
			}
		}
		if twos == 2 {
			return 2 // two of a kind
		}
		panic(fmt.Sprintf("not expecting %v", combined))
	}
	if len(combined) == 4 {
		return 1 // pair
	}
	if len(combined) == 5 {
		return 0 //high card
	}
	return -1
}

func (c ComparerPartTwo) handType(hand string) int {
	combined := map[rune]int{}
	jackCount := 0
	highestRune := '0'
	highestCount := -1
	for _, r := range hand {
		if r == 'J' {
			jackCount++
			continue
		}
		v, f := combined[r]
		if f {
			combined[r]++
		} else {
			combined[r] = 1
		}
		if v+1 > highestCount {
			highestCount = v + 1
			highestRune = r
		}
	}
	combined[highestRune] += jackCount

	if len(combined) == 1 {
		return 6 // five of a kind
	}
	if len(combined) == 2 {
		for _, v := range combined {
			if v == 1 || v == 4 {
				return 5 // four of a kind
			} else {
				return 4 // full house
			}
		}
	}
	if len(combined) == 3 {
		twos := 0
		for _, v := range combined {
			if v == 2 {
				twos++
			}
			if v == 3 {
				return 3 // three of a kind
			}
		}
		if twos == 2 {
			return 2 // two of a kind
		}
		panic(fmt.Sprintf("not expecting %v", combined))
	}
	if len(combined) == 4 {
		return 1 // pair
	}
	if len(combined) == 5 {
		return 0 //high card
	}
	return -1
}

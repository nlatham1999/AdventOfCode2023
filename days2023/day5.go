package days2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayFive() {
	// RunPartOne()
	RunPartTwo()
}

func RunPartTwo() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day5.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	seeds := strings.Split(inputs[0], " ")[1:]

	// seedMap := map[int][]int{}
	splits := [][]int{}
	for i, _ := range seeds {
		if i%2 == 0 {
			start, _ := strconv.Atoi(seeds[i])
			numRange, _ := strconv.Atoi(seeds[i+1])
			splits = append(splits, []int{start, numRange, 0})
		}
	}

	skip := false
	round := 0
	firstRow := ""
	key := ""
	nextKey := ""
	tree := make(map[string][][]int)
	nextTree := make(map[string]string)
	for _, row := range inputs[1:] {
		if row == "" {
			skip = true
			continue
		}
		if skip {
			skip = false
			round++

			t := strings.Split(row, " ")
			t = strings.Split(t[0], "-")

			key = t[0]
			nextKey = t[2]

			tree[key] = [][]int{}
			nextTree[key] = nextKey

			continue
		}
		firstRow = row

		nums1 := strings.Split(firstRow, " ")
		d1, err := strconv.Atoi(nums1[0])
		if err != nil {
			fmt.Println("Could not parse")
		}
		s1, err := strconv.Atoi(nums1[1])
		if err != nil {
			fmt.Println("Could not parse")
		}
		r1, err := strconv.Atoi(nums1[2])
		if err != nil {
			fmt.Println("Could not parse")
		}

		tree[key] = append(tree[key], []int{d1, s1, r1})

		firstRow = ""
	}

	min := -1

	key = "seed"
	continueRunning := true
	stopNext := false

	for _, s := range splits {
		fmt.Println(s)
	}

	for continueRunning {
		fmt.Println(fmt.Sprintf("going into %s-to-%s", key, nextTree[key]))

		newSplits := [][]int{}
		vals := tree[key]

		valsSorted := []int{}
		valTree := map[int][]int{}
		for _, x := range vals {
			valsSorted = append(valsSorted, x[1])
			valTree[x[1]] = x
		}
		valsSorted = internal.SortIntAsc(valsSorted)

		splitsSorted := []int{}
		splitTree := map[int][]int{}
		for _, x := range splits {
			splitsSorted = append(splitsSorted, x[0])
			splitTree[x[0]] = x
		}
		splitsSorted = internal.SortIntAsc(splitsSorted)

		i := 0
		j := 0
		// fmt.Println(fmt.Sprintf("Splits Sorted %v", splits))
		sStart := splitTree[splitsSorted[i]][0]
		count := 0
		for i < len(splitsSorted) && j < len(valsSorted) {
			s := splitTree[splitsSorted[i]]
			v := valTree[valsSorted[j]]
			sEnd := s[0] + s[1]
			vEnd := v[1] + v[2]

			count++
			if count == 1000 {
				fmt.Println("Hit count max")
				return
			}
			// fmt.Println(count)

			// fmt.Println(fmt.Sprintf("%d, %d, %d, %v, %v", i, j, sStart, s, v))

			if sStart < v[1] {
				if sEnd < v[1] {
					newSplits = append(newSplits, []int{sStart, s[1], s[2]})
					i++
					if i < len(splitsSorted) {
						sStart = splitTree[splitsSorted[i]][0]
					}
				} else {
					newSplits = append(newSplits, []int{sStart, v[1] - sStart, s[2]})
					sStart = v[1]
				}
			} else if sStart >= v[1] && sStart < vEnd {
				// fmt.Println("test 2")
				offset := v[0] - v[1]
				if sEnd <= vEnd {
					newSplits = append(newSplits, []int{sStart, s[1] - (sStart - s[0]), s[2] + offset})
					i++
					if i < len(splitsSorted) {
						sStart = splitTree[splitsSorted[i]][0]
					}
				} else {
					newSplits = append(newSplits, []int{sStart, vEnd - sStart, s[2] + offset})
					sStart = vEnd
				}
			} else if v[1] < sStart {
				offset := v[0] - v[1]
				if sStart < vEnd {
					if sEnd <= vEnd {
						newSplits = append(newSplits, []int{sStart, s[1], s[2] + offset})
						i++
						if i < len(splitsSorted) {
							sStart = splitTree[splitsSorted[i]][0]
						}
					} else {
						newSplits = append(newSplits, []int{sStart, vEnd - sStart, s[2] + offset})
						j++
						sStart = vEnd
					}
				} else {
					j++
				}
			}

		}

		for i < len(splitsSorted) {
			s := splitTree[splitsSorted[i]]
			newSplits = append(newSplits, []int{sStart, s[0] + s[1] - sStart, s[2]})
			i++
			if i < len(splitsSorted) {
				sStart = splitTree[splitsSorted[i]][0]
			}
		}

		splits = newSplits

		key = nextTree[key]

		// fmt.Println(val)
		// fmt.Println("Next Key " + key)
		if stopNext {
			continueRunning = false
		}
		if key == "location" {
			stopNext = true
		}

		for _, s := range splits {
			fmt.Println(s)
		}
		fmt.Println(fmt.Sprintf("COUNT %d", count))
	}

	fmt.Println("Splits")
	for _, s := range splits {
		if min == -1 || s[0]+s[2] < min {
			min = s[0] + s[2]
		}
	}

	fmt.Println(min)

}

func RunPartOne() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day5.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	seeds := strings.Split(inputs[0], " ")[1:]
	fmt.Println(seeds)

	seedMap := map[int][]int{}
	for _, seed := range seeds {
		i, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("could not parse")
			return
		}
		seedMap[i] = []int{}
	}

	skip := false
	round := 0
	firstRow := ""
	key := ""
	nextKey := ""
	tree := make(map[string][][]int)
	nextTree := make(map[string]string)
	for _, row := range inputs[1:] {
		if row == "" {
			skip = true
			continue
		}
		if skip {
			skip = false
			round++

			t := strings.Split(row, " ")
			t = strings.Split(t[0], "-")

			key = t[0]
			nextKey = t[2]

			tree[key] = [][]int{}
			nextTree[key] = nextKey

			continue
		}
		firstRow = row

		nums1 := strings.Split(firstRow, " ")
		d1, err := strconv.Atoi(nums1[0])
		if err != nil {
			fmt.Println("Could not parse")
		}
		s1, err := strconv.Atoi(nums1[1])
		if err != nil {
			fmt.Println("Could not parse")
		}
		r1, err := strconv.Atoi(nums1[2])
		if err != nil {
			fmt.Println("Could not parse")
		}

		tree[key] = append(tree[key], []int{d1, s1, r1})

		firstRow = ""
	}

	min := -1

	for seed, _ := range seedMap {
		val := seed
		key = "seed"
		continueRunning := true
		stopNext := false
		for continueRunning {
			fmt.Println(fmt.Sprintf("%d going into %s-to-%s", val, key, nextTree[key]))
			vals := tree[key]
			val = XtoY(vals, val)

			key = nextTree[key]
			fmt.Println(val)
			fmt.Println("Next Key " + key)
			if stopNext {
				continueRunning = false
			}
			if key == "location" {
				stopNext = true
			}
		}
		if min == -1 || val < min {
			min = val
		}
	}

	fmt.Println(min)

}

func XtoY(vals [][]int, seed int) int {
	for _, values := range vals {
		// fmt.Println(fmt.Sprintf("Range one %d to %d", values[1], values[1]+values[2]))
		if seed >= values[1] && seed <= values[1]+values[2] {
			diff := seed - values[1]
			return values[0] + diff
		}
	}
	return seed
}

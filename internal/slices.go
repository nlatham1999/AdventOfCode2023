package internal

import "sort"

func SumOfSliceInt(arr []int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func SortIntDesc(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	return numbers
}

func SortIntAsc(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	return numbers
}

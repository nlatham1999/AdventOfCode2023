package days2023

import (
	"fmt"
	"math"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDay11() {
	// RunDay11PartOne()
	RunDay11PartTwo()
}

type XYPair struct {
	X int
	Y int
}

func RunDay11PartTwo() {

	amountToIncrease := 1000000

	inputs, err := internal.ReadFileLineByLine("./inputs2023/day11.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	//expand the galaxy
	galaxyRune := [][]rune{}
	for _, row := range inputs {
		galaxyRune = append(galaxyRune, []rune(row))
	}

	clearColumns := make(map[int]interface{})
	clearRows := make(map[int]interface{})

	for i := range galaxyRune[0] {
		// if it needs to be expanded
		clear := columnClear(i, galaxyRune)
		if clear {
			clearColumns[i] = nil
		}
	}

	for i, _ := range galaxyRune {
		clear := rowClear(i, galaxyRune)
		if clear {
			clearRows[i] = nil
		}
	}

	pairs := []XYPair{}
	for i, row := range galaxyRune {
		for j, x := range row {
			if x == '#' {
				pairs = append(pairs, XYPair{X: j, Y: i})
			}
		}
	}

	sum := 0.0
	for i, p1 := range pairs {
		for j, p2 := range pairs {
			if i != j {
				countX := 0
				k := math.Min(float64(p1.X), float64(p2.X))
				for k < math.Max(float64(p1.X), float64(p2.X)) {
					kint := int(k)
					if _, f := clearColumns[kint]; f {
						countX += amountToIncrease
					} else {
						countX += 1
					}
					k++
				}

				countY := 0
				k = math.Min(float64(p1.Y), float64(p2.Y))
				for k < math.Max(float64(p1.Y), float64(p2.Y)) {
					kint := int(k)
					if _, f := clearRows[kint]; f {
						countY += amountToIncrease
					} else {
						countY += 1
					}
					k++
				}

				sum += (float64(countX) + float64(countY)) / 2
			}
		}
	}

	fmt.Printf("\n%f \n", sum)

}

func RunDay11PartOne() {

	inputs, err := internal.ReadFileLineByLine("./inputs2023/day11.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	//expand the galaxy
	galaxyRune := [][]rune{}
	newGalaxy := [][]rune{}
	newNewGalaxy := [][]rune{}
	for _, row := range inputs {
		galaxyRune = append(galaxyRune, []rune(row))
		newGalaxy = append(newGalaxy, []rune{})
	}

	for i := range galaxyRune[0] {
		// if it needs to be expanded
		clear := columnClear(i, galaxyRune)
		for j, v := range galaxyRune {
			newGalaxy[j] = append(newGalaxy[j], v[i])
			if clear {
				newGalaxy[j] = append(newGalaxy[j], v[i])
			}
		}
	}

	for i, v := range newGalaxy {
		clear := rowClear(i, newGalaxy)
		newNewGalaxy = append(newNewGalaxy, v)
		if clear {
			newNewGalaxy = append(newNewGalaxy, v)
		}
	}

	for _, v := range newNewGalaxy {
		fmt.Println(string(v))
	}

	pairs := []XYPair{}
	for i, row := range newNewGalaxy {
		for j, x := range row {
			if x == '#' {
				pairs = append(pairs, XYPair{X: j, Y: i})
			}
		}
	}

	sum := 0.0
	for i, p1 := range pairs {
		for j, p2 := range pairs {
			if i != j {
				sum += math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y))
			}
		}
	}

	fmt.Printf("\n%f \n", sum/2)

}

func rowClear(y int, r [][]rune) bool {
	for _, v := range r[y] {
		if v != '.' {
			return false
		}
	}
	return true
}

func columnClear(x int, r [][]rune) bool {
	for _, v := range r {
		if v[x] != '.' {
			return false
		}
	}
	return true
}

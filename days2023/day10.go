package days2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDay10() {
	occurences := RunDay10PartOne()
	RunDay10PartTwo(occurences)
}

func RunDay10PartTwo(occurences [][]rune) {

	inputs, err := internal.ReadFileLineByLine("./inputs2023/day10.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	matrix := [][]rune{}

	sReplacement := '-' // this changes based on input

	for i, row := range inputs {
		rRow := []rune(row)
		// matrix = append(matrix, rRow)
		newRow := []rune{}
		rowBelow := []rune{}
		for j, r := range rRow {
			if occurences[i][j] != 'Y' {
				r = '.'
			}
			if r == 'S' {
				r = sReplacement
			}
			newRow = append(newRow, r)
			if r == '-' || (r == 'F') || r == 'L' {
				newRow = append(newRow, r)
			} else {
				newRow = append(newRow, ',')
			}
			if r == '|' || r == 'F' || r == '7' {
				rowBelow = append(rowBelow, r)
			} else {
				rowBelow = append(rowBelow, ',')
			}
			rowBelow = append(rowBelow, ',')
		}
		matrix = append(matrix, newRow)
		matrix = append(matrix, rowBelow)
	}

	for _, row := range matrix {
		for _, x := range row {
			fmt.Print(string(x))
		}
		fmt.Println()
	} 

	count := 0
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if matrix[i][j] == ',' || matrix[i][j] == '.' {
				x := j
				y := i
				start := strconv.Itoa(x) + "_" + strconv.Itoa(y)
				queue := make(map[string]interface{})
				queue[start] = nil
				rcount := 0
				invalid := false
				for len(queue) > 0 {

					topKey := ""
					for k := range queue {
						topKey = k
						break
					}
					key := strings.Split(topKey, "_")
					delete(queue, topKey)

					x = internal.StringToIntFast(key[0])
					y = internal.StringToIntFast(key[1])

					if matrix[y][x] != ',' && matrix[y][x] != '.' {
						continue
					}

					if matrix[y][x] == '.' {
						rcount++
					}

					matrix[y][x] = ' '

					if y > 0 {
						if matrix[y-1][x] == ',' || matrix[y-1][x] == '.' {
							k := strconv.Itoa(x) + "_" + strconv.Itoa(y-1)
							queue[k] = nil
						}
					}
					if y < len(matrix)-1 {
						if matrix[y+1][x] == ',' || matrix[y+1][x] == '.' {
							k := strconv.Itoa(x) + "_" + strconv.Itoa(y+1)
							queue[k] = nil
						}
					}
					if x > 0 {
						if matrix[y][x-1] == ',' || matrix[y][x-1] == '.' {
							k := strconv.Itoa(x-1) + "_" + strconv.Itoa(y)
							queue[k] = nil
						}
					}
					if x < len(matrix[y])-1 {
						if matrix[y][x+1] == ',' || matrix[y][x+1] == '.' {
							k := strconv.Itoa(x+1) + "_" + strconv.Itoa(y)
							queue[k] = nil
						}
					}

					if x == 0 || y == 0 || x == len(matrix[y])-1 || y == len(matrix)-1 {
						invalid = true
					}
				}

				if !invalid {
					count += rcount
				}
			}
		}
	}

	// for _, row := range matrix {
	// 	for _, x := range row {
	// 		fmt.Print(string(x))
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(count)
}

func RunDay10PartOne() [][]rune {

	inputs, err := internal.ReadFileLineByLine("./inputs2023/day10.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	matrix := [][]rune{}
	partOfLine := [][]rune{}

	x := 0
	y := 0
	for i, row := range inputs {
		rRow := []rune(row)
		rRow2 := []rune(row)
		matrix = append(matrix, rRow)
		partOfLine = append(partOfLine, rRow2)
		for j, r := range rRow {
			if r == 'S' {
				y = i
				x = j
			}
		}
	}

	hitS := false
	count := 0

	fmt.Println(matrix)

	previousY := -1
	previousX := -1
	for !hitS {

		// fmt.Println(fmt.Sprintf("x: %d, y: %d, prevX: %d, prevY: %d, val: %s", x, y, previousX, previousY, string(matrix[y][x])))

		partOfLine[y][x] = 'Y'
		if matrix[y][x] == 'S' {
			previousX = x
			previousY = y
			if y > 0 && (matrix[y-1][x] == 'F' || matrix[y-1][x] == '7' || matrix[y-1][x] == '|') {
				y = y - 1
			} else if y < len(matrix)-1 && (matrix[y+1][x] == 'J' || matrix[y+1][x] == 'L' || matrix[y+1][x] == '|') {
				y = y + 1
			} else if x > 0 && (matrix[y][x-1] == 'L' || matrix[y][x-1] == 'F' || matrix[y][x-1] == '-') {
				x = x - 1
			} else if x < len(matrix[x])-1 && (matrix[y][x+1] == 'J' || matrix[y][x+1] == '7' || matrix[y][x+1] == '-') {
				x = x + 1
			}
		} else if matrix[y][x] == '|' {
			if previousX == x && previousY == y-1 {
				previousY = y
				y = y + 1
			} else {
				previousY = y
				y = y - 1
			}
		} else if matrix[y][x] == 'F' {
			if previousX == x && previousY == y+1 {
				previousY = y
				previousX = x
				x = x + 1
			} else {
				previousY = y
				previousX = x
				y = y + 1
			}
		} else if matrix[y][x] == '7' {
			if previousX == x && previousY == y+1 {
				previousY = y
				previousX = x
				x = x - 1
			} else {
				previousY = y
				previousX = x
				y = y + 1
			}
		} else if matrix[y][x] == 'J' {
			if previousX == x && previousY == y-1 {
				previousY = y
				previousX = x
				x = x - 1
			} else {
				previousY = y
				previousX = x
				y = y - 1
			}
		} else if matrix[y][x] == 'L' {
			if previousX == x && previousY == y-1 {
				previousY = y
				previousX = x
				x = x + 1
			} else {
				previousY = y
				previousX = x
				y = y - 1
			}
		} else if matrix[y][x] == '-' {
			if previousX == x-1 && previousY == y {
				previousY = y
				previousX = x
				x = x + 1
			} else {
				previousY = y
				previousX = x
				x = x - 1
			}
		} else {
			fmt.Println("UNKNOWN")
		}

		if matrix[y][x] == 'S' {
			hitS = true
		}

		count++
	}

	print(count)

	return partOfLine

}

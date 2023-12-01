package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileLineByLine(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ReadFileAsInts(filepath string) []int {
	inputs, err := ReadFileLineByLine(filepath)
	if err != nil {
		fmt.Println("COuld not read file")
		return nil
	}

	ints := []int{}
	for _, x := range inputs {
		i, err := strconv.Atoi(x)
		if err != nil {
			continue
		}
		ints = append(ints, i)
	}
	return ints
}

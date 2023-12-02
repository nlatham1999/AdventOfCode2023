package days2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlatham1999/AdventOfCode2023/internal"
)

func RunDayTwo() {

	DayTwoPartOne()
	DayTwoPartTwo()
}

func DayTwoPartOne() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day2.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	sum := 0
	for _, game := range inputs {
		sum += parseGames(game)
	}
	fmt.Println(sum)
}

func DayTwoPartTwo() {
	inputs, err := internal.ReadFileLineByLine("./inputs2023/day2.txt")
	if err != nil {
		fmt.Println("Could not open")
	}

	sum := 0
	for _, game := range inputs {
		sum += parseGames2(game)
	}
	fmt.Println(sum)
}

func parseGames(game string) int {
	gameText := strings.Split(game, ":")
	rounds := strings.Split(gameText[1], ";")
	for _, round := range rounds {
		pieces := strings.Split(round, ",")
		for _, piece := range pieces {
			pieceTrimmed := strings.TrimSpace(piece)
			s := strings.Split(pieceTrimmed, " ")
			num, err := strconv.Atoi(s[0])
			color := s[1]
			if err != nil {
				fmt.Println(fmt.Sprintf("Got an error trying to parse the int. piece: %v", s[0]))
			}
			if color == "red" && num > 12 {
				return 0
			}
			if color == "green" && num > 13 {
				return 0
			}
			if color == "blue" && num > 14 {
				return 0
			}
		}
	}
	gameId, err := strconv.Atoi(strings.Split(gameText[0], " ")[1])
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to get game id. %v", gameText[0]))
	}
	return gameId
}

func parseGames2(game string) int {
	gameText := strings.Split(game, ":")
	rounds := strings.Split(gameText[1], ";")
	numRed := 0
	numGreen := 0
	numBlue := 0
	for _, round := range rounds {
		pieces := strings.Split(round, ",")
		for _, piece := range pieces {
			pieceTrimmed := strings.TrimSpace(piece)
			s := strings.Split(pieceTrimmed, " ")
			num, err := strconv.Atoi(s[0])
			color := s[1]
			if err != nil {
				fmt.Println(fmt.Sprintf("Got an error trying to parse the int. piece: %v", s[0]))
			}
			if color == "red" {
				if num > numRed {
					numRed = num
				}
			}
			if color == "green" {
				if num > numGreen {
					numGreen = num
				}
			}
			if color == "blue" {
				if num > numBlue {
					numBlue = num
				}
			}
		}
	}
	return numRed * numGreen * numBlue
}

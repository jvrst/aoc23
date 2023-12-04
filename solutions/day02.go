package solutions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	Red   string = "r"
	Green string = "g"
	Blue  string = "b"
)

func SolveDay02() {
    lines := GetLines(2)

	var validGames []int
	var gamePowers []int

    for _, line := range lines {
		// var val int
		nums := fndNums(line)
		gameNum := getGameNr(line)
		maxNums := getMaxColors(nums)

		// Part 1
		if isValidGame(maxNums) {
			validGames = append(validGames, gameNum)
		}
		// Part 2
		gamePower := multiply(maxNums)
		gamePowers = append(gamePowers, gamePower)

		log.Println(gameNum)
	}

	fmt.Println("Part 1: Valid games: ", validGames)
	fmt.Println(sum(validGames))
	fmt.Println("Part 2: Game powers: ", gamePowers)
	fmt.Println(sum(gamePowers))
}

func multiply(maxColors []int) int {
	return maxColors[0] * maxColors[1] * maxColors[2]
}

func getMaxColors(colors [][]string) []int {
	curMax := []int{0, 0, 0}
	for _, color := range colors {
		s := strings.Split(color[0], " ")
		numstr, val := s[0], s[1]
		num, err := strconv.Atoi(numstr)
		if err != nil {
			log.Fatal(err)
		}
		switch val {
		case Red:
			if curMax[0] < num {
				curMax[0] = num
			}
		case Green:
			if curMax[1] < num {
				curMax[1] = num
			}
		case Blue:
			if curMax[2] < num {
				curMax[2] = num
			}
		default:
			log.Println("Error")
		}
	}
	fmt.Print("CurMax: ", curMax)
	return curMax
}

func fndNums(text string) [][]string {
	var re = regexp.MustCompile(`\d+ \w`)
	matches := re.FindAllStringSubmatch(text, -1)
	log.Println(text, matches)
	return matches
}

func getGameNr(text string) int {
	var re = regexp.MustCompile(`\d+`)
	match := re.FindString(text)
	gameNum, err := strconv.Atoi(match)
	if err != nil {
		log.Fatal(err)
	}
	return gameNum
}

func isValidGame(maxColors []int) bool {
	if maxColors[0] <= 12 && maxColors[1] <= 13 && maxColors[2] <= 14 {
		return true
	}
	return false
}

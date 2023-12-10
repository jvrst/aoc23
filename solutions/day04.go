package solutions

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func SolveDay04() {
	lines := GetLines(4)

	SolveDay04Part1(lines)
	SolveDay04Part2(lines)
}

type Card struct {
	id          int
	winningNums []int
	haveNums    []int
}

func SolveDay04Part2(lines []string) {
	cards := parseCards(lines)
	fmt.Println(cards)
	var cardScores []int
	for _, card := range cards {
		winningNumbers := getWinningNumbers(card)
		cardScore := calculateScore(winningNumbers)

		cardScores = append(cardScores, cardScore)
	}
	fmt.Println(sum(cardScores))
}

func SolveDay04Part1(lines []string) {
	cards := parseCards(lines)
	fmt.Println(cards)
	var cardScores []int
	for _, card := range cards {
		winningNumbers := getWinningNumbers(card)
		cardScore := calculateScore(winningNumbers)

		cardScores = append(cardScores, cardScore)
	}
	fmt.Println(sum(cardScores))
}

func parseCards(lines []string) []Card {
	cards := []Card{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		splitGame := strings.Split(line, ":")
		game := strings.Split(splitGame[0], " ")
		gameNumbers := strings.Split(splitGame[1], "|")
		winningNumbersStr := re.FindAllString(gameNumbers[0], -1)
		haveNumsStr := re.FindAllString(gameNumbers[1], -1)

		var winningNumbers []int
		var haveNums []int
		for _, winningNum := range winningNumbersStr {
			num, _ := strconv.Atoi(winningNum)
			winningNumbers = append(winningNumbers, num)
		}
		for _, haveNum := range haveNumsStr {
			num, _ := strconv.Atoi(haveNum)
			haveNums = append(haveNums, num)
		}
		gameNr, _ := strconv.Atoi(game[1])
		card := Card{
			id:          gameNr,
			winningNums: winningNumbers,
			haveNums:    haveNums,
		}
		cards = append(cards, card)
	}
	return cards
}

func getWinningNumbers(card Card) []int {
	var winningNumbers []int
	for _, haveNum := range card.haveNums {
		for _, winNum := range card.winningNums {
			if haveNum == winNum {
				winningNumbers = append(winningNumbers, haveNum)
			}
		}
	}
	return winningNumbers
}

func calculateScore(winningNumbers []int) int {
	return powInt(2, len(winningNumbers)-1)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

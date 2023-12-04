package solutions

import (
	"fmt"
	"log"

	"regexp"
	"strconv"
)

func findStringNums(text string, last bool) (int, error) {
	var re *regexp.Regexp
	if last {
		re = regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|\d)`)
	} else {
		re = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	}
	matches := re.FindAllStringSubmatch(text, -1)
	if len(matches) == 0 {
		return -1, fmt.Errorf("No numbers")
	}
	var nums []int

	for _, match := range matches {
		var num int
		switch match[1] {
		case "one":
			num = 1
		case "two":
			num = 2
		case "three":
			num = 3
		case "four":
			num = 4
		case "five":
			num = 5
		case "six":
			num = 6
		case "seven":
			num = 7
		case "eight":
			num = 8
		case "nine":
			num = 9
		default:
			var err error
			num, err = strconv.Atoi(match[1])
			if err != nil {
				return -1, err
			}
		}
		nums = append(nums, num)
	}
	if len(nums) == 1 {
		nums = append(nums, nums[0])
	}
	fmt.Println(text, nums)
	return nums[0], nil
}

func findNums(text string) ([]int, error) {
	re := regexp.MustCompile(`\b(\d)\b`)
	matches := re.FindAllStringSubmatch(text, -1)
	var nums []int
	for _, match := range matches {
		fmt.Println(match)
		num, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func SolveDay01() {
    lines := GetLines(1)

	var calibrationValues []int

    for _, line := range lines {
		// var val int
		first, _ := findStringNums(line, false)
		last, _ := findStringNums(line, true)
		num := first*10 + last
		fmt.Println(num)

		calibrationValues = append(calibrationValues, num)

	}
	fmt.Println(calibrationValues)
	fmt.Println(sum(calibrationValues))

}

func sum(nrs []int) int {
	total := 0
	for _, v := range nrs {
		total += v
	}
	return total
}

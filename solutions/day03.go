package solutions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Coordinates struct {
	x int
	y int
}

type Gear struct {
	coordinates Coordinates
	parts       []PartNumber
}

type PartNumber struct {
	coordinates []Coordinates
	number      int
	valid       bool
}

func SolveDay03() {
	lines := GetLines(3)

	// Add padding to all lines
	for k, line := range lines {
		lines[k] = "." + line + "."
	}
	padLine := strings.Repeat(".", len(lines[1]))
	lines = append([]string{padLine}, lines...)
	lines = append(lines, padLine)

	// Create a 2D array of runes
	schematic := make([][]rune, 0)
	for _, line := range lines {
		schematic = append(schematic, []rune(line))
	}

	var partNumbers []PartNumber
	for k, line := range lines {
		partNumbers = append(partNumbers, getPartNumbers(line, k)...)
	}

	var valids int
	for k, partNum := range partNumbers {
		for _, coord := range partNum.coordinates {
			if hasNeighbourSymbols(schematic, coord) {
				partNumbers[k].valid = true
			}
		}
		if partNumbers[k].valid {
			valids += partNum.number
		}
	}

	// Part 2: Process gears
	var validGears []Gear
	gearLocs := findGears(schematic)
	for _, gearLoc := range gearLocs {
		neighbourCoords := hasNeighbourNumbers(schematic, gearLoc)

		// Use a map to ensure unique part numbers per gear
		uniqueParts := make(map[int]PartNumber)
		for _, coord := range neighbourCoords {
			for _, partNum := range partNumbers {
				for _, partCoord := range partNum.coordinates {
					if coord.x == partCoord.x && coord.y == partCoord.y {
						uniqueParts[partNum.number] = partNum
					}
				}
			}
		}

		// Create a gear only if exactly two unique parts are found
		if len(uniqueParts) == 2 {
			var gearParts []PartNumber
			for _, part := range uniqueParts {
				gearParts = append(gearParts, part)
			}
			fullGear := Gear{gearLoc, gearParts}
			validGears = append(validGears, fullGear)
		}
	}

	for _, gear := range validGears {
		fmt.Println(gear)
	}
	fmt.Printf("Part 1: Sum of all part numbers: %d\n", valids)
	ratioSum := 0
	for _, gear := range validGears {
		ratio := gear.parts[0].number * gear.parts[1].number
		ratioSum += ratio
	}
	fmt.Printf("Part 2: Sum of all gear ratios: %d\n", ratioSum)
}

func isSymbol(r rune) bool {
	return r != '.' && !isDigit(r)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func getPartNumbers(line string, lineNr int) []PartNumber {
	re := regexp.MustCompile(`\d+`)
	matchIndices := re.FindAllStringIndex(string(line), -1)
	matches := re.FindAllString(string(line), -1)
	fmt.Println(line, matches, matchIndices)
	var nums []PartNumber
	for k, match := range matches {
		// Parse the actual number
		num, err := strconv.Atoi(match)
		if err != nil {
			log.Fatal(err)
		}

		// Get coordinates for each digit
		var coords []Coordinates
		for i := matchIndices[k][0]; i < matchIndices[k][1]; i++ {
			coords = append(coords, Coordinates{x: lineNr, y: i})
		}

		// Create struct with the number + all the relevant coordinates
		partNum := PartNumber{
			coordinates: coords,
			number:      num,
		}
		nums = append(nums, partNum)
	}
	return nums
}

func hasNeighbourSymbols(schematic [][]rune, coords Coordinates) bool {
	x := coords.x
	y := coords.y
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy
			if isSymbol(schematic[nx][ny]) {
				fmt.Println("Found symbol: ", nx, ny, string(schematic[nx][ny]))
				return true
			}
		}
	}
	return false
}

func findGears(schematic [][]rune) []Coordinates {
	var gears []Coordinates
	for x, line := range schematic {
		for y, r := range line {
			if r == '*' {
				gears = append(gears, Coordinates{x: x, y: y})
			}
		}
	}
	return gears
}

func hasNeighbourNumbers(schematic [][]rune, coords Coordinates) []Coordinates {
	x := coords.x
	y := coords.y
	neighbourCoords := []Coordinates{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy
			if isDigit(schematic[nx][ny]) {
				// fmt.Println("Found symbol: ", nx, ny, string(schematic[nx][ny]))
				neighbourCoords = append(neighbourCoords, Coordinates{x: nx, y: ny})
			}
		}
	}
	return neighbourCoords
}

package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetFile(dayNr int) *os.File {
	day := fmt.Sprintf("%02d", dayNr)
	path := "./files/day" + day + ".txt"
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func GetLines(dayNr int) []string {
	day := fmt.Sprintf("%02d", dayNr)
	path := "./files/day" + day + ".txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

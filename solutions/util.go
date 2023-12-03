package solutions

import (
	"fmt"
	"log"
	"os"
)

func GetFile(dayNr int) *os.File {
	day := fmt.Sprintf("%02d", dayNr)
	file, err := os.Open("./files/day" + day + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return file
}

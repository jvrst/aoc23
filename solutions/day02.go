package solutions

import (
  "os"
  "log"

  "regexp"
)

func SolveDay02() {
	file, err := os.Open("./files/day02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}

func fndNums(text string) {
  var re = regexp.MustCompile(`\b(\d \w)\b`)
  re.FindAllStringSubmatchIndex(text, -1)
  log.Println(re.FindAllStringSubmatchIndex(text, -1))
}

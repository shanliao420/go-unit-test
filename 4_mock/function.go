package __mock

import (
	"bufio"
	"log"
	"os"
)

func ReadFirstLine() string {
	open, err := os.Open("log.log")
	defer open.Close()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	targetLine := line + "😎"
	return targetLine
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var leftToRightParentheses = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func GetPenaltyCorruptedLines() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//line := scanner.Text()

	}

	return 0
}

func main() {
	fmt.Println(GetPenaltyCorruptedLines())
}

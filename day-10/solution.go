package main

import (
	"bufio"
	"container/list"
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

func findFirstInvalidCharacter(line string) string {
	stack := list.New()
	for _, char := range line {
		if val, ok := leftToRightParentheses[string(char)]; ok {
			stack.PushBack(val)
		} else if stack.Len() == 0 {
			return string(char)
		} else {
			par := stack.Back()
			if par.Value != leftToRightParentheses[string(char)] {
				return string(char)
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(GetPenaltyCorruptedLines())
}

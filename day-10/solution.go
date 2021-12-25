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

var characterToValue = map[string]int{
	"":  0,
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
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

	var errorScore int
	for scanner.Scan() {
		character := findFirstInvalidCharacter(scanner.Text())
		errorScore += characterToValue[character]
	}

	return errorScore
}

func findFirstInvalidCharacter(line string) string {
	stack := Stack{}
	for _, char := range line {
		par := string(char)
		if _, ok := leftToRightParentheses[par]; ok {
			stack.Push(par)
		} else if len(stack) == 0 {
			return par
		} else {
			leftPar := stack.Pop()
			if par != leftToRightParentheses[leftPar] {
				return par
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(GetPenaltyCorruptedLines())
}

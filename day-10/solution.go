package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var leftToRightParentheses = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var corruptedCharacterToValue = map[string]int{
	"":  0,
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var missingCharacterToValue = map[string]int{
	"":  0,
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
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

	var penaltyScore int
	for scanner.Scan() {
		character := findFirstInvalidCharacter(scanner.Text())
		penaltyScore += corruptedCharacterToValue[character]
	}

	return penaltyScore
}

func GetPenaltyIncompleteLines() int {
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

	var penaltyScores []int
	for scanner.Scan() {
		linePenalty := 0
		characters := findMissingCharacters(scanner.Text())

		for i := range characters {
			linePenalty *= 5
			linePenalty += missingCharacterToValue[characters[len(characters)-1-i]]
		}

		if linePenalty != 0 {
			penaltyScores = append(penaltyScores, linePenalty)
		}
	}

	sort.Ints(penaltyScores)
	return penaltyScores[len(penaltyScores)/2]
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

func findMissingCharacters(line string) []string {
	stack := Stack{}
	for _, char := range line {
		par := string(char)
		if _, ok := leftToRightParentheses[par]; ok {
			stack.Push(par)
		} else if len(stack) == 0 {
			return []string{}
		} else {
			leftPar := stack.Pop()
			if par != leftToRightParentheses[leftPar] {
				return []string{}
			}
		}
	}
	return stack
}

func main() {
	fmt.Println(GetPenaltyCorruptedLines())
	fmt.Println(GetPenaltyIncompleteLines())
}

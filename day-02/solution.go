package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func unpackInput(inp string) (string, int) {
	s := strings.Split(inp, " ")
	return s[0], convertStrToInt(s[1])
}

func convertStrToInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return converted
}

func FindSubmarinePosition() int {
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

	horizontalPos, verticalPost := 0, 0

	for scanner.Scan() {
		action, value := unpackInput(scanner.Text())
		if action == "forward" {
			horizontalPos += value
		} else if action == "down" {
			verticalPost += value
		} else if action == "up" {
			verticalPost -= value
		}
	}

	return horizontalPos * verticalPost
}

func FindSubmarinePositionWithAim() int {
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

	horizontalPos, verticalPost, aim := 0, 0, 0

	for scanner.Scan() {
		action, value := unpackInput(scanner.Text())
		if action == "forward" {
			horizontalPos += value
			verticalPost += aim * value
		} else if action == "down" {
			aim += value
		} else if action == "up" {
			aim -= value
		}
	}

	return horizontalPos * verticalPost
}

func main() {
	fmt.Println(FindSubmarinePosition())
	fmt.Println(FindSubmarinePositionWithAim())
}

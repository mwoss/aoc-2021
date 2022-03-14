package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func convertStrToInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return converted
}

func main() {
	rawContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(rawContent)

	reg, err := regexp.Compile("([-]?\\d+..[-]?\\d+)")
	if err != nil {
		log.Fatal(err)
	}

	matches := reg.FindAllString(content, -1)
	if len(matches) != 2 {
		log.Fatal("Input file should only contains x and y ranges")
	}

	x := strings.Split(matches[0], "..")
	y := strings.Split(matches[1], "..")

	x1, x2 := convertStrToInt(x[0]), convertStrToInt(x[1])
	y1, y2 := convertStrToInt(y[0]), convertStrToInt(y[1])

	fmt.Println(x1, x2)
	fmt.Println(y1, y2)

	// x - initial x velocity value
	// y - initial y velocity value
	// (x1..x2) = x + (x-1) + (x-2) + (x-3) + ... (if positive)
	// (x1..x2) = x + (x+1) + (x+2) + (x+3) + ... (if negative)
	// (y1..y2) = y + (y-1) + (y-2) + (y-3) + ... (all cases)
	n := -y1 - 1
	fmt.Println(n * (n + 1) / 2) // height decrease by 1 with each step, so maybe we could just calculate the sequence of n..1 numbers?
}
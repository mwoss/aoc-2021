package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countDepthMeasurementIncrease() int {
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
	scanner.Scan() // move to the first line of file

	increased, prev := 0, convertStrToInt(scanner.Text())
	for scanner.Scan() {
		curr := convertStrToInt(scanner.Text())
		if prev < curr {
			increased++
		}
		prev = curr
	}
	return increased
}

func main() {
	fmt.Println(countDepthMeasurementIncrease())
}

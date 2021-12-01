package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readMeasurement(scanner *bufio.Scanner) int {
	scanner.Scan()
	return convertStrToInt(scanner.Text())
}

func countDepthMeasurementIncreaseSlidingWindow() int {
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

	increased := 0
	first, second, third := readMeasurement(scanner), readMeasurement(scanner), readMeasurement(scanner)
	for scanner.Scan() {
		curr := convertStrToInt(scanner.Text())
		if first + second + third < second + third + curr {
			increased++
		}
		first, second, third = second, third, curr
	}
	return increased
}

func main() {
	fmt.Println(countDepthMeasurementIncreaseSlidingWindow())
}

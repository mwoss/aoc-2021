package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func convertStrToInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return converted
}

func readMeasurement(scanner *bufio.Scanner) int {
	scanner.Scan()
	return convertStrToInt(scanner.Text())
}

func CountDepthMeasurementIncrease() int {
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

func CountDepthMeasurementIncreaseSlidingWindow() int {
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
		if first+second+third < second+third+curr {
			increased++
		}
		first, second, third = second, third, curr
	}
	return increased
}

func main() {
	fmt.Println(CountDepthMeasurementIncrease())
	fmt.Println(CountDepthMeasurementIncreaseSlidingWindow())
}

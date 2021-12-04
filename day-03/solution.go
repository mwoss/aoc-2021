package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func bufferToNumericGamma(buffer []int) (int64, error) {
	var strBin string
	for _, value := range buffer {
		if value > 0 {
			strBin += "1"
		} else {
			strBin += "0"
		}
	}
	return strconv.ParseInt(strBin, 2, 32)
}

func bufferToNumericEpsilon(buffer []int) (int64, error) {
	var strBin string
	for _, value := range buffer {
		if value > 0 {
			strBin += "0"
		} else {
			strBin += "1"
		}
	}
	return strconv.ParseInt(strBin, 2, 32)
}

func GetSubmarinePowerConsumption() int {
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

	gammaBuffer := make([]int, 12) // we assume all values have the same length

	for scanner.Scan() {
		measurement := scanner.Text()
		for pos, digit := range measurement {
			if string(digit) == "1" {
				gammaBuffer[pos]++
			} else {
				gammaBuffer[pos]--
			}
		}
	}

	gamma, err := bufferToNumericGamma(gammaBuffer)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := bufferToNumericEpsilon(gammaBuffer)
	if err != nil {
		log.Fatal(err)
	}

	return int(gamma * epsilon)
}

func GetLifeSupportRating() int {
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

	zerosPrefixedTemp := make([]string, 0)
	onesPrefixedTemp := make([]string, 0)

	// initial data read
	for scanner.Scan() {
		measurement := scanner.Text()
		if string(measurement[0]) == "0" {
			zerosPrefixedTemp = append(zerosPrefixedTemp, measurement)
		} else {
			onesPrefixedTemp = append(onesPrefixedTemp, measurement)
		}
	}

	// let's assume the input si always correct
	var mostCommonValue, leastCommonValue string
	mostComparator := func(a []string, b []string) bool { return len(a) <= len(b) }
	leastComparator := func(a []string, b []string) bool { return len(a) > len(b) }

	// could be separated to sub function, but it's not a production code xD
	if len(zerosPrefixedTemp) <= len(onesPrefixedTemp) {
		mostCommonValue = divideMeasurementsByOccurrenceComparator(onesPrefixedTemp, mostComparator)
		leastCommonValue = divideMeasurementsByOccurrenceComparator(zerosPrefixedTemp, leastComparator)
	} else {
		mostCommonValue = divideMeasurementsByOccurrenceComparator(zerosPrefixedTemp, mostComparator)
		leastCommonValue = divideMeasurementsByOccurrenceComparator(onesPrefixedTemp, leastComparator)
	}

	oxygenRating, _ := strconv.ParseInt(mostCommonValue, 2, 32)
	co2Rating, _ := strconv.ParseInt(leastCommonValue, 2, 32)

	return int(oxygenRating * co2Rating)
}

func divideMeasurementsByOccurrenceComparator(measurements []string, comparator func([]string, []string) bool) string {
	posIdx := 1
	for len(measurements) != 1 {
		zerosPrefixed := make([]string, 0)
		onesPrefixed := make([]string, 0)

		for _, measurement := range measurements {
			if string(measurement[posIdx]) == "0" {
				zerosPrefixed = append(zerosPrefixed, measurement)
			} else {
				onesPrefixed = append(onesPrefixed, measurement)
			}
		}

		if comparator(zerosPrefixed, onesPrefixed) {
			measurements = onesPrefixed
		} else {
			measurements = zerosPrefixed
		}

		posIdx++
	}

	return measurements[0]
}

func main() {
	fmt.Println(GetSubmarinePowerConsumption())
	fmt.Println(GetLifeSupportRating())
}

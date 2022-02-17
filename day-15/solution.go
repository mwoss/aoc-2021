package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x int
	y int
}

func readRiskLevels(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var riskLevels [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riskLevels = append(riskLevels, unpackInputLine(scanner.Text()))
	}

	if len(riskLevels) == 0 {
		log.Fatal("risk level grid has been incorrectly parsed, collection length should be greater than 0")
	}

	return riskLevels
}

func unpackInputLine(inp string) []int {
	var levels []int
	for _, level := range strings.Split(inp, "") {
		levels = append(levels, convertStrToInt(level))
	}
	return levels
}

func convertStrToInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return converted
}

func calculateCurrentLevel(initialRisk, offset int) int {
	return (initialRisk+offset)%10 + (initialRisk+offset)/10
}

func FindLowestRiskLevelPath(riskLevels [][]int) int {
	nodeCount := len(riskLevels) * len(riskLevels[0])
	rowCount := len(riskLevels[0])

	visited := make([]bool, nodeCount)
	distances := make([]int, nodeCount)

	for i := range distances {
		distances[i] = math.MaxInt32 // infinity in terms of risk level
	}

	distances[0] = 0

	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: 0, priority: 0})

	for pq.Len() != 0 {
		node := heap.Pop(&pq).(*Item)

		if visited[node.value] {
			continue
		}

		visited[node.value] = true

		if distances[node.value] < node.priority {
			continue
		}

		rowId := node.value / rowCount
		colId := node.value - (rowId * rowCount)
		potentialNeighbors := []Node{
			{rowId + 1, colId},
			{rowId - 1, colId},
			{rowId, colId + 1},
			{rowId, colId - 1},
		}

		for _, neighbor := range potentialNeighbors {
			if neighbor.x < 0 || neighbor.x >= len(riskLevels) || neighbor.y < 0 || neighbor.y >= rowCount {
				continue
			}

			nodeId := neighbor.x*len(riskLevels) + neighbor.y
			newDistance := -distances[node.value] + riskLevels[neighbor.x][neighbor.y]

			if newDistance < distances[nodeId] {
				distances[nodeId] = -newDistance
				heap.Push(&pq, &Item{value: nodeId, priority: -newDistance})
			}
		}

		if node.value == nodeCount-1 {
			return -distances[nodeCount-1]
		}
	}

	return math.MaxInt32
}

func FindLowestRiskLevelPathOnExtendedMap(riskLevels [][]int) int {
	initHeight, initWidth := len(riskLevels), len(riskLevels[0])
	for i := range riskLevels {
		for j := 1; j < 5; j++ {
			for k := 0; k < initWidth; k++ {
				riskLevels[i] = append(riskLevels[i], calculateCurrentLevel(riskLevels[i][k], j))
			}
		}
	}
	for i := 1; i < 5; i++ {
		for j := 0; j < initHeight; j++ {
			var extRow []int
			for k := 0; k < len(riskLevels[j]); k++ {
				extRow = append(extRow, calculateCurrentLevel(riskLevels[j][k], i))
			}
			riskLevels = append(riskLevels, extRow)
		}
	}
	return FindLowestRiskLevelPath(riskLevels)
}

func main() {
	riskLevels := readRiskLevels("input.txt")
	fmt.Println(FindLowestRiskLevelPath(riskLevels))
	fmt.Println(FindLowestRiskLevelPathOnExtendedMap(riskLevels))
}

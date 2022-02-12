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

func unpackInput(inp string) []int {
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

func FindLowestRiskLevel() int {
	file, err := os.Open("input2.txt")
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
		riskLevels = append(riskLevels, unpackInput(scanner.Text()))
	}

	if len(riskLevels) == 0 {
		log.Fatal("risk level grid has been incorrectly parsed, collection length should be greater than 0")
	}

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

func main() {
	fmt.Println(FindLowestRiskLevel())
}

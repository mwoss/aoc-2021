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

	var riskLevels [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riskLevels = append(riskLevels, unpackInput(scanner.Text()))
	}

	nodeCount := len(riskLevels) * len(riskLevels[0])
	rowCount := len(riskLevels[0])

	visited := make([]bool, nodeCount)
	dist := make([]int, nodeCount)

	for i := range dist {
		dist[i] = math.MaxInt32 // infinity in terms of risk level
	}

	dist[0] = 0

	pq := PriorityQueue{}
	heap.Init(&pq)

	heap.Push(&pq, &Item{value: 0, priority: 0})

	for pq.Len() != 0 {
		node := heap.Pop(&pq).(*Item)

		visited[node.value] = true

		if dist[node.value] < node.priority {
			continue
		}

		potentialEdges := []int{node.value + rowCount, node.value - rowCount, node.value + 1, node.value - 1}
		for _, edge := range potentialEdges {
			if edge <= 0 || edge >= nodeCount {
				continue
			}
			if visited[edge] {
				continue
			}

			edgeRow := edge / rowCount
			edgeCol := edge - (edgeRow * rowCount)

			newDist := dist[node.value] + riskLevels[edgeCol][edgeCol] // fix

			if newDist < dist[edge] {
				dist[edge] = newDist
				heap.Push(&pq, &Item{value: edge, priority: newDist})
			}
		}

		if node.value == nodeCount-1 {
			return dist[nodeCount-1]
		}
	}

	return math.MaxInt32
}

func main() {
	fmt.Println(FindLowestRiskLevel())
}

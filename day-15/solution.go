package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
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

type Node struct {
	index int
	dist  int
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
		dist[i] = 10 // infinity in terms of risk level
	}

	dist[0] = 0

	pq := list.New()
	pq.PushFront(Node{0, 0})

	for pq.Len() != 0 {
		nodeEl := pq.Front()
		pq.Remove(nodeEl)

		node := nodeEl.Value.(Node)

		visited[node.index] = true

		if dist[node.index] < node.dist {
			continue
		}

		potentialEdges := []int{node.index + rowCount, node.index - rowCount, node.index + 1, node.index - 1}

		for _, edge := range potentialEdges {
			if edge <= 0 || edge >= nodeCount {
				continue
			}
			if visited[edge] {
				continue
			}

			edgeRow := edge / rowCount
			edgeCol := edge - (edgeRow * rowCount)

			newDist := dist[node.index] + riskLevels[edgeCol][edgeCol] // fix

			if newDist < dist[edge] {
				dist[edge] = newDist
				pq.PushFront(Node{edge, newDist})
			}
		}

		if node.index == nodeCount-1 {
			return dist[nodeCount]
		}
	}

	return 10
}

func main() {
	fmt.Println(FindLowestRiskLevel())
}

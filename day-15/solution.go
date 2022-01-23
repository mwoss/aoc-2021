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

	fmt.Println(riskLevels)

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
		node := pq.Front().Value.(Node)

		visited[node.index] = true

		if dist[node.index] < node.dist {
			continue
		}


	}

	return 0
}

func main() {
	fmt.Println(FindLowestRiskLevel())
}

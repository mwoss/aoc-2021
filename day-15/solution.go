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

// ............
// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Node struct {
	x int
	y int
}

///...............

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

		edgeRow := node.value / rowCount
		edgeCol := node.value - (edgeRow * rowCount)

		potentialEdges := []Node{{edgeRow + 1, edgeCol}, {edgeRow - 1, edgeCol}, {edgeRow, edgeCol + 1}, {edgeRow, edgeCol - 1}}
		for _, edge := range potentialEdges {
			if edge.x < 0 || edge.x >= len(riskLevels) || edge.y < 0 || edge.y >= rowCount {
				continue
			}

			nodeId := edge.x*len(riskLevels) + edge.y

			if visited[nodeId] {
				continue
			}

			//edgeRow := edge / rowCount
			//edgeCol := edge - (edgeRow * rowCount)

			newDist := dist[node.value] + riskLevels[edge.x][edge.y] // fix

			if newDist < dist[nodeId] {
				dist[nodeId] = newDist
				heap.Push(&pq, &Item{value: nodeId, priority: newDist})
			}
		}

		if node.value == nodeCount-1 {
			fmt.Println(dist[:10])
			return dist[nodeCount-1]
		}
	}

	return math.MaxInt32
}

func main() {
	fmt.Println(FindLowestRiskLevel())
}

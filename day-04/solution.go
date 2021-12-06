package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type NumberPosition struct {
	boardId int
	x       int
	y       int
}

type BingoCell struct {
	number   string
	isMarked bool
}

func prepareBingoBoards(unparsedBoards []string) ([][5][5]BingoCell, map[string][]NumberPosition) {
	var boards [][5][5]BingoCell
	numberToPositions := make(map[string][]NumberPosition)

	for boardId, unparsedBoard := range unparsedBoards {
		board := [5][5]BingoCell{}
		for rowId, row := range strings.Split(unparsedBoard, "\n") {
			for colId, num := range strings.Fields(row) {
				board[rowId][colId] = BingoCell{num, false}
				numberToPositions[num] = append(numberToPositions[num], NumberPosition{boardId, rowId, colId})
			}
		}
		boards = append(boards, board)
	}

	return boards, numberToPositions
}

func getSumOfUnmarkedNumbers(board *[5][5]BingoCell) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board[i][j].isMarked {
				n, _ := strconv.Atoi(board[i][j].number)
				sum += n
			}
		}
	}
	return sum
}

func checkRow(board *[5][5]BingoCell, row int) bool {
	for i := 0; i < 5; i++ {
		if !board[row][i].isMarked {
			return false
		}
	}
	return true
}

func checkCol(board *[5][5]BingoCell, col int) bool {
	for i := 0; i < 5; i++ {
		if !board[i][col].isMarked {
			return false
		}
	}
	return true
}

func FindFirstWinningBoardScore() int {
	b, err := ioutil.ReadFile("day-04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Split(string(b), "\n\n")
	numbers, unparsedBoards := strings.Split(content[0], ","), content[1:]

	boards, numberToPositions := prepareBingoBoards(unparsedBoards)

	for _, number := range numbers {
		positions := numberToPositions[number]
		for _, position := range positions {
			boards[position.boardId][position.x][position.y].isMarked = true
			if checkRow(&boards[position.boardId], position.x) || checkCol(&boards[position.boardId], position.y) {
				winNumber, _ := strconv.Atoi(number)
				return getSumOfUnmarkedNumbers(&boards[position.boardId]) * winNumber
			}
		}
	}

	return -1
}

func FindLastWinningBoardScore() int {
	b, err := ioutil.ReadFile("day-04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Split(string(b), "\n\n")
	numbers, unparsedBoards := strings.Split(content[0], ","), content[1:]

	boards, numberToPositions := prepareBingoBoards(unparsedBoards)

	recentlyWonScore := -1
	alreadyWonBoards := make(map[int]struct{})

	for _, number := range numbers {
		positions := numberToPositions[number]
		for _, position := range positions {
			if _, ok := alreadyWonBoards[position.boardId]; ok {
				// if board was already marked as "won" we don't and shouldn't update its state
				// otherwise we would end up with incorrect end solution
				continue
			}

			boards[position.boardId][position.x][position.y].isMarked = true
			if checkRow(&boards[position.boardId], position.x) || checkCol(&boards[position.boardId], position.y) {
				winNumber, _ := strconv.Atoi(number)
				recentlyWonScore = getSumOfUnmarkedNumbers(&boards[position.boardId]) * winNumber
				alreadyWonBoards[position.boardId] = struct{}{}
			}
		}
	}

	return recentlyWonScore
}

func main() {
	fmt.Println(FindFirstWinningBoardScore())
	fmt.Println(FindLastWinningBoardScore())
}

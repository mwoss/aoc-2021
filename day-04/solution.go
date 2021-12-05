package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func prepareBingoBoards(scanner *bufio.Scanner) ([][5][5]BingoCell, map[string][]NumberPosition) {
	var boards [][5][5]BingoCell
	board, rowId := [5][5]BingoCell{}, 0
	numberToPositions := make(map[string][]NumberPosition)

	for scanner.Scan() {
		values := scanner.Text()
		if values == "" {
			boards = append(boards, board)
			board, rowId = [5][5]BingoCell{}, 0
			continue
		}

		for i, num := range strings.Fields(values) {
			board[rowId][i] = BingoCell{num, false}
			numberToPositions[num] = append(numberToPositions[num], NumberPosition{len(boards), rowId, i})
		}
		rowId++
	}
	boards = append(boards, board)

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
	file, err := os.Open("day-04/input.txt")
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
	scanner.Scan()

	numbers := strings.Split(scanner.Text(), ",")

	scanner.Scan() // skip empty line, I will learn how to do it better in future aoc problems :3

	boards, numberToPositions := prepareBingoBoards(scanner)

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
	file, err := os.Open("day-04/input.txt")
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
	scanner.Scan()

	numbers := strings.Split(scanner.Text(), ",")

	scanner.Scan() // skip empty line, I will learn how to do it better in future aoc problems :3

	boards, numberToPositions := prepareBingoBoards(scanner)

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

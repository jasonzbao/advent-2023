package main

import (
	"fmt"
	"utils"
)

var dummy = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func isReflection(board [][]string, line float64, isVertical bool, success func(int) bool) bool {
	// flip the board
	if !isVertical {
		newBoard := make([][]string, len(board[0]))
		for i := 0; i < len(board[0]); i++ {
			newBoard[i] = make([]string, len(board))
		}
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				newBoard[j][i] = board[i][j]
			}
		}
		board = newBoard
	}
	errors := 0
	i := .5
	for {
		lineLeft, lineRight := int(line-i), int(line+i)
		if lineLeft < 0 || lineRight >= len(board) {
			return success(errors)
		}
		for j := 0; j < len(board[lineLeft]); j++ {
			if board[lineLeft][j] != board[lineRight][j] {
				errors += 1
			}
		}
		i += 1
	}
}

func part1(boards [][][]string) {
	sum := 0
	for boardNumber, board := range boards {
		fmt.Println(board)
		for i := .5; i < float64(len(board[0])-1); i++ {
			if isReflection(board, i, false, func(errors int) bool {
				return errors == 0
			}) {
				fmt.Println("board", boardNumber, "is reflection on ", i+.5, "vertical")
				sum += int(i + .5)
			}
		}
		for i := .5; i < float64(len(board)-1); i++ {
			if isReflection(board, i, true, func(errors int) bool {
				return errors == 0
			}) {
				fmt.Println("board", boardNumber, "is reflection on ", i+.5, "horizontal")
				sum += int(i+.5) * 100
			}
		}
	}
	fmt.Println(sum)
}

func part2(boards [][][]string) {
	sum := 0
	for boardNumber, board := range boards {
		fmt.Println(board)
		for i := .5; i < float64(len(board[0])-1); i++ {
			if isReflection(board, i, false, func(errors int) bool {
				return errors == 1
			}) {
				fmt.Println("board", boardNumber, "is reflection on ", i+.5, "vertical")
				sum += int(i + .5)
			}
		}
		for i := .5; i < float64(len(board)-1); i++ {
			if isReflection(board, i, true, func(errors int) bool {
				return errors == 1
			}) {
				fmt.Println("board", boardNumber, "is reflection on ", i+.5, "horizontal")
				sum += int(i+.5) * 100
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	input := utils.FormattedRequest(13)
	//input = utils.FormatInput(dummy)

	boards := make([][][]string, 0)
	board := make([][]string, 0)
	i := 0
	for _, line := range input {
		if line == "" {
			boards = append(boards, board)
			board = make([][]string, 0)
			i = 0
			continue
		}
		board = append(board, make([]string, 0))
		for _, char := range line {
			board[i] = append(board[i], string(char))
		}
		i++
	}
	boards = append(boards, board)

	part1(boards)
	part2(boards)
}

package main

import (
	"fmt"
	"utils"
)

const dummy = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func touchAround(x, y int, input []string, fn func(int, int)) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i < 0 || j < 0 || i >= len(input) || j >= len(input[i]) {
				continue
			}
			fn(i, j)
		}
	}
}

func part1(input []string) {
	sum := 0
	seen := make(map[int]bool)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if utils.IsDigit(input[i][j]) || string(input[i][j]) == "." {
				continue
			}
			touchAround(i, j, input, func(x, y int) {
				if seen[x*len(input[y])+y] {
					return
				}
				if utils.IsDigit(input[x][y]) {
					partNumber := utils.ExtractNum(y, input[x], seen, x)
					sum += partNumber
				}
			})

		}
	}
	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if string(input[i][j]) != "*" {
				continue
			}
			gears := make([]int, 0)
			seen := make(map[int]bool)
			touchAround(i, j, input, func(x, y int) {
				if seen[x*len(input[y])+y] {
					return
				}
				if utils.IsDigit(input[x][y]) {
					partNumber := utils.ExtractNum(y, input[x], seen, x)
					gears = append(gears, partNumber)
				}
			})
			if len(gears) == 2 {
				sum += gears[0] * gears[1]
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	input := utils.FormattedRequest(3)
	part1(input)
	part2(input)
}

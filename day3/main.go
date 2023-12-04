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

func part1(input []string) {
	sum := 0
	seen := make(map[int]bool)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if utils.IsDigit(input[i][j]) || string(input[i][j]) == "." {
				continue
			}
			utils.TouchAround(i, j, input, func(x, y int) {
				if seen[x*len(input[y])+y] {
					return
				}
				if utils.IsDigit(input[x][y]) {
					partNumber, start, end := utils.ExtractNum(y, input[x])
					for k := start; k <= end; k++ {
						seen[x*len(input[y])+k] = true
					}
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
			utils.TouchAround(i, j, input, func(x, y int) {
				if seen[x*len(input[y])+y] {
					return
				}
				if utils.IsDigit(input[x][y]) {
					partNumber, start, end := utils.ExtractNum(y, input[x])
					for k := start; k <= end; k++ {
						seen[x*len(input[y])+k] = true
					}
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

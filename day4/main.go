package main

import (
	"fmt"
	"strings"
	"utils"
)

const dummy = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func part1(lines []*utils.Line) {
	total := 0
	for _, line := range lines {
		split := strings.Split(line.Line, "|")
		winners := make(map[string]bool)
		for _, s := range strings.Split(strings.TrimSpace(split[0]), " ") {
			if len(s) == 0 {
				continue
			}
			winners[strings.TrimSpace(s)] = true
			fmt.Println(s)
		}
		winnerCount := 0
		for _, s := range strings.Split(strings.TrimSpace(split[1]), " ") {
			if winners[strings.TrimSpace(s)] {
				winnerCount++
				fmt.Println("Winner", s, winnerCount)
			}
		}
		if winnerCount > 0 {
			total += utils.GenericPow[int](2, winnerCount-1)
		}
	}
	fmt.Println(total)
}

func part2(lines []*utils.Line) {
	multipliers := make(map[int]int)
	for _, line := range lines {
		split := strings.Split(line.Line, "|")
		winners := make(map[string]bool)
		for _, s := range strings.Split(strings.TrimSpace(split[0]), " ") {
			if len(s) == 0 {
				continue
			}
			winners[strings.TrimSpace(s)] = true
		}
		winnerCount := 0
		for _, s := range strings.Split(strings.TrimSpace(split[1]), " ") {
			if winners[strings.TrimSpace(s)] {
				winnerCount++
			}
		}
		// default 1, if already set, add 1 for the initial
		if _, ok := multipliers[line.Index]; !ok {
			multipliers[line.Index] = 1
		} else {
			multipliers[line.Index]++
		}
		for i := line.Index + 1; i <= line.Index+winnerCount; i++ {
			multipliers[i] += multipliers[line.Index]
		}
	}
	total := 0
	for _, v := range multipliers {
		total += v
	}
	fmt.Println(total)
}

func main() {
	input := utils.FormattedRequest(4)
	lines := utils.ParseFormat(input)
	part1(lines)
	part2(lines)
}

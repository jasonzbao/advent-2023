package main

import (
	"fmt"
	"unicode"
	"utils"
)

func part1(input []string) {
	total := 0
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				total += int(line[i]-'0') * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				total += int(line[i] - '0')
				break
			}
		}
	}
	fmt.Println(total)
}

var nums = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseInput(s *string, i int, rev bool) int {
	if unicode.IsDigit(rune((*s)[i])) {
		return int((*s)[i] - '0')
	}
	for k, v := range nums {
		var substr string
		if rev {
			if i-len(k)+1 < 0 {
				continue
			}
			substr = (*s)[i-len(k)+1 : i+1]
		} else {
			if i+len(k) > len(*s) {
				continue
			}
			substr = (*s)[i : i+len(k)]
		}
		if substr == k {
			return v
		}
	}
	return 0
}

func part2(input []string) {
	total := 0
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			parsed := parseInput(&line, i, false)
			total += parsed * 10
			if parsed > 0 {
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			parsed := parseInput(&line, i, true)
			total += parsed
			if parsed > 0 {
				break
			}
		}
	}
	fmt.Println(total)
}

func main() {
	input := utils.FormattedRequest(1)
	part1(input)
	part2(input)
}

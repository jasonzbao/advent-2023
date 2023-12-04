package utils

import (
	"log"
	"strconv"
)

func ParseRune(s uint8) int {
	if s >= '0' && s <= '9' {
		return int(s - '0')
	}
	return -1
}

func IsDigit(s uint8) bool {
	return ParseRune(s) != -1
}

func ExtractNum(i int, s string) (val, start, end int) {
	var endIndex, startIndex int
	for j := i; j < len(s); j++ {
		if IsDigit(s[j]) {
			endIndex = j
		} else {
			break
		}
	}

	for j := i; j >= 0; j-- {
		if IsDigit(s[j]) {
			startIndex = j
		} else {
			break
		}
	}

	ret, err := strconv.Atoi(s[startIndex : endIndex+1])
	if err != nil {
		log.Fatal("Error parsing number", err)
	}
	return ret, startIndex, endIndex
}

func TouchAround(x, y int, input []string, fn func(int, int)) {
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

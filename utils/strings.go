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

func ExtractNum(i int, s string, seen map[int]bool, z int) int {
	var endIndex, startIndex int
	for j := i; j < len(s); j++ {
		if IsDigit(s[j]) {
			endIndex = j
		} else {
			break
		}
		seen[z*len(s)+j] = true
	}

	for j := i; j >= 0; j-- {
		if IsDigit(s[j]) {
			startIndex = j
		} else {
			break
		}
		seen[z*len(s)+j] = true
	}

	ret, err := strconv.Atoi(s[startIndex : endIndex+1])
	if err != nil {
		log.Fatal("Error parsing number", err)
	}
	return ret
}

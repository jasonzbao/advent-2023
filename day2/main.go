package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

type Token struct {
	Color string
	Num   int
}

func tokenize(s string) [][]Token {
	var err error
	exp := strings.Split(s, ";")
	tokens := make([][]Token, len(exp))
	for i := 0; i < len(exp); i++ {
		balls := strings.Split(exp[i], ",")
		tokens[i] = make([]Token, len(balls))
		for j := 0; j < len(balls); j++ {
			balls[j] = strings.TrimSpace(balls[j])
			splitBalls := strings.Split(balls[j], " ")
			tokens[i][j].Num, err = strconv.Atoi(splitBalls[0])
			if err != nil {
				log.Panic("Error parsing balls", err)
			}
			tokens[i][j].Color = splitBalls[1]
		}
	}
	return tokens
}

var part1Map = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(in [][][]Token) {
	sum := 0
	for i, line := range in {
		sum += i + 1
		gameSet := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, grab := range line {
			for _, token := range grab {
				if token.Num > part1Map[token.Color] {
					gameSet[token.Color] += token.Num
				}
			}
		}
		for k, v := range part1Map {
			if gameSet[k] > v {
				sum -= i + 1
				break
			}
		}
	}

	fmt.Println(sum)
}

func part2(in [][][]Token) {
	sum := 0
	for _, line := range in {
		gameSet := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, grab := range line {
			for _, token := range grab {
				gameSet[token.Color] = utils.GenericMax[int](gameSet[token.Color], token.Num)
			}
		}
		sum += gameSet["red"] * gameSet["green"] * gameSet["blue"]
	}
	fmt.Println(sum)
}

func main() {
	input := utils.FormattedRequest(2)

	in := make([][][]Token, 0)
	for _, line := range input {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Panic("Invalid parts", parts)
		}
		in = append(in, tokenize(parts[1]))
	}

	part1(in)
	part2(in)
}

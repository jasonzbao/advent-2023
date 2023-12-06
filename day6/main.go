package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"utils"
)

const dummy = `Time:      7  15   30
Distance:  9  40  200`

// not ideal but to simulate > instead of >=
const wiggle = .0000001

func parse(input []string) []int {
	parsed := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if i == 0 {
			continue
		}
		s := strings.TrimSpace(input[i])
		if s == "" {
			continue
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		parsed = append(parsed, val)
	}
	return parsed
}

func numberWins(time, distance int) int {
	q := utils.GenericPow[int](time, 2) - 4*distance
	if q < 0 {
		return 0
	}
	if time%2 == 0 {
		return int(math.Floor(math.Sqrt(float64(q))/2-wiggle)*2 + 1)
	} else {
		return int(math.Floor((math.Sqrt(float64(q))/2-wiggle)+.5) * 2)
	}
}

func part1(times []int, distances []int) {
	ret := 1
	for i := 0; i < len(times); i++ {
		n := numberWins(times[i], distances[i])
		if n > 0 {
			ret *= n
		}
	}
	fmt.Println(ret)
}

func part2(time, distance int) {
	fmt.Println(numberWins(time, distance))
}

func main() {
	input := utils.FormattedRequest(6)
	//input := utils.FormatInput(dummy)
	times := strings.Split(input[0], " ")
	distances := strings.Split(input[1], " ")
	part1(parse(times), parse(distances))
	part2(49877895, 356137815021882)
}

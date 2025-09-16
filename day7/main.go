// Note, the answer here is to part 2 where it treats J as joker
package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type Hand struct {
	Cards string
	Bid   int
}

// 0 = high card, 1 = one pair, 2 = two pair, 3 = three of a kind, 4 = full house,
func returnRank(cards string) int {
	occurences := make(map[string]int)
	jokers := 0
	for _, card := range cards {
		if string(card) == "J" {
			jokers++
		} else {
			occurences[string(card)]++
		}
	}

	values := []int{}
	for _, v := range occurences {
		values = append(values, v)
	}

	sort.Ints(values)

	if jokers > 0 {
		if len(values) == 0 {
			return 6
		}
		values[len(values)-1] += jokers
	}

	// 5 = four of a kind, 6 = five of a kind
	if values[len(values)-1] == 5 || values[len(values)-1] == 4 {
		return values[len(values)-1] + 1
	}

	// 4 = full house
	if values[len(values)-1] == 3 && values[len(values)-2] == 2 {
		return 4
	}

	// 3 = three of a kind
	if values[len(values)-1] == 3 {
		return 3
	}

	// 2 = two pair
	if values[len(values)-1] == 2 && values[len(values)-2] == 2 {
		return 2
	}

	// 1 = one pair
	if values[len(values)-1] == 2 {
		return 1
	}

	return 0
}

func convertCard(card string) int {
	if card == "A" {
		return 14
	}
	if card == "K" {
		return 13
	}
	if card == "Q" {
		return 12
	}
	if card == "T" {
		return 10
	}
	if card == "J" {
		return -1
	}
	return int(card[0] - '0')
}

func tieBreak(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return convertCard(string(a[i])) > convertCard(string(b[i]))
		}
	}
	panic("Tie break failed")
}

func isLarger(a, b Hand) bool {
	rankA := returnRank(a.Cards)
	rankB := returnRank(b.Cards)

	if rankA != rankB {
		return rankA > rankB
	}
	return tieBreak(a.Cards, b.Cards)
}

func part1(input []string) {
	hands := make([]Hand, 0)
	for _, line := range input {
		parts := strings.Split(line, " ")

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		hands = append(hands, Hand{
			Cards: parts[0],
			Bid:   bid,
		})
	}

	sort.Slice(hands, func(j, i int) bool {
		return isLarger(hands[i], hands[j])
	})

	res := 0
	for i := 0; i < len(hands); i++ {
		fmt.Println(hands[i].Cards, hands[i].Bid, i+1)
		res += hands[i].Bid * (i + 1)
	}
	fmt.Println(res)
}

func main() {
	input := utils.FormattedRequest(7)

	// input := utils.FormatInput(`32T3K 765
	// T55J5 684
	// KK677 28
	// KTJJT 220
	// QQQJA 483`)

	part1(input)
}

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	HighCard CardType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type CardType int

var cardValues = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Hand struct {
	CardType
	cards string
	bid   int
}

func parseHand(s string) Hand {
	r := regexp.MustCompile(`^(.*) (.*)$`)
	m := r.FindStringSubmatch(s)
	cards := m[1]
	bid, _ := strconv.Atoi(m[2])
	h := Hand{getCardType(cards), cards, bid}
	return h
}

func (h Hand) Compare(hh Hand) int {
	if h.CardType < hh.CardType {
		return -1
	}
	if h.CardType > hh.CardType {
		return 1
	}
	for i := range h.cards {
		a := cardValues[h.cards[i]]
		b := cardValues[hh.cards[i]]
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
	}
	return 0
}

func (h Hand) Jokerized() HandWithJoker {
	hh := HandWithJoker{Hand{getCardTypeWithJoker(h.cards), h.cards, h.bid}}
	return hh
}

func getCardType(cards string) CardType {
	m := make(map[byte]int)
	for _, r := range []byte(cards) {
		if _, ok := m[r]; !ok {
			m[r] = 0
		}
		m[r] += 1
	}
	max := 0
	pairCount := 0
	for _, count := range m {
		if count > max {
			max = count
		}
		if count == 2 {
			pairCount += 1
		}
	}
	switch {
	case max == 5:
		return FiveOfAKind
	case max == 4:
		return FourOfAKind
	case max == 3 && pairCount == 1:
		return FullHouse
	case max == 3:
		return ThreeOfAKind
	case pairCount == 2:
		return TwoPair
	case pairCount == 1:
		return OnePair
	default:
		return HighCard
	}
}

type HandWithJoker struct {
	Hand
}

func (h HandWithJoker) Compare(hh HandWithJoker) int {
	if h.CardType < hh.CardType {
		return -1
	}
	if h.CardType > hh.CardType {
		return 1
	}
	for i := range h.cards {
		a := cardValues[h.cards[i]]
		if a == 11 {
			a = 0
		}

		b := cardValues[hh.cards[i]]
		if b == 11 {
			b = 0
		}

		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
	}
	return 0
}

func getCardTypeWithJoker(cards string) CardType {
	m := make(map[byte]int)
	for _, r := range []byte(cards) {
		if _, ok := m[r]; !ok {
			m[r] = 0
		}
		m[r] += 1
	}
	joker := 0
	max := 0
	pairCount := 0
	for c, count := range m {
		if c == 'J' {
			joker = count
			continue
		}
		if count > max {
			max = count
		}
		if count == 2 {
			pairCount += 1
		}
	}
	switch {
	case max+joker == 5:
		return FiveOfAKind
	case max+joker == 4:
		return FourOfAKind
	case (max == 3 && pairCount == 1) || (pairCount == 2 && joker == 1):
		return FullHouse
	case max+joker == 3:
		return ThreeOfAKind
	case pairCount == 2:
		return TwoPair
	case pairCount == 1 || max+joker == 2:
		return OnePair
	default:
		return HighCard
	}
}

type Solver struct {
	hands []Hand
}

func NewSolver(input string) *Solver {
	hands := []Hand{}
	for _, line := range strings.Split(input, "\n") {
		hands = append(hands, parseHand(line))
	}
	s := Solver{hands}
	return &s
}

func (s *Solver) Part1() int {
	hands := make([]Hand, len(s.hands))
	copy(hands, s.hands)
	less := func(i, j int) bool {
		c := hands[i].Compare(hands[j])
		return c < 0
	}
	sort.Slice(hands, less)
	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	return sum
}

func (s *Solver) Part2() int {
	hands := []HandWithJoker{}
	for _, hand := range s.hands {
		hands = append(hands, hand.Jokerized())
	}
	less := func(i, j int) bool {
		c := hands[i].Compare(hands[j])
		return c < 0
	}
	sort.Slice(hands, less)
	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}

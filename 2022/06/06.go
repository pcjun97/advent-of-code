package main

import (
	"fmt"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	fmt.Println(IndexStartOfPacketMarker(input) + 1)
	fmt.Println(IndexStartOfMessageMarker(input) + 1)
}

func indexMarker(s string, count int) int {
	set := make(map[byte]int)

	var i int
	for i = 0; i < len(s); i++ {
		c := s[i]

		if _, ok := set[c]; ok {
			set[c]++
		} else {
			set[c] = 1
		}

		if i >= count {
			c = s[i-count]
			if set[c] == 1 {
				delete(set, c)
			} else {
				set[c]--
			}
		}

		if len(set) == count {
			return i
		}
	}

	return -1
}

func IndexStartOfPacketMarker(s string) int {
	return indexMarker(s, 4)
}

func IndexStartOfMessageMarker(s string) int {
	return indexMarker(s, 14)
}

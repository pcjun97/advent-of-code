package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	DoorID []byte
}

func NewSolver(input string) *Solver {
	s := Solver{[]byte(input)}
	return &s
}

func (s *Solver) md5(salt int) [16]byte {
	hash := md5.Sum(append(s.DoorID, []byte(strconv.Itoa(salt))...))
	return hash
}

func (s *Solver) Part1() string {
	pw := ""
	salt := 0

	for {
		hash := s.md5(salt)
		salt++

		if !isHashStartsWith00000(hash) {
			continue
		}

		pw = fmt.Sprintf("%s%x", pw, hash[2])
		if len(pw) >= 8 {
			break
		}
	}

	return pw
}

func (s *Solver) Part2() string {
	pm := make(map[byte]byte)
	salt := 0

	for {
		hash := s.md5(salt)
		salt++

		if !isHashStartsWith00000(hash) {
			continue
		}

		p := hash[2]
		if p >= 8 {
			continue
		}

		if _, ok := pm[p]; ok {
			continue
		}

		pm[p] = (hash[3] >> 4)

		if len(pm) >= 8 {
			break
		}
	}

	pw := ""
	for i := 0; i < 8; i++ {
		pw = fmt.Sprintf("%s%x", pw, pm[byte(i)])
	}

	return pw
}

func isHashStartsWith00000(hash [16]byte) bool {
	return (hash[0] | hash[1] | (hash[2] & 0xf0)) == 0
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}

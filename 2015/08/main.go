package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	lines []string
}

func NewSolver(input string) *Solver {
	s := Solver{
		lines: strings.Split(input, "\n"),
	}
	return &s
}

func (s *Solver) Part1() int {
	sumLiteral := 0
	sumMemory := 0

	for _, line := range s.lines {
		sumLiteral += len(line)
		sumMemory += len([]rune(parseString(line)))
	}

	return sumLiteral - sumMemory
}

func (s *Solver) Part2() int {
	sumOriginal := 0
	sumEncoded := 0

	for _, line := range s.lines {
		sumOriginal += len(line)
		sumEncoded += len(encodeString(line))
	}

	return sumEncoded - sumOriginal
}

func parseString(s string) string {
	r := regexp.MustCompile(`^"(.*)"$`)
	m := r.FindStringSubmatch(s)

	source := m[1]
	result := ""
	for i := 0; i < len(source); i++ {
		if source[i] != '\\' {
			result = result + source[i:i+1]
			continue
		}

		switch source[i+1] {
		case '\\':
			result += "\\"
			i += 1
			continue
		case '"':
			result += "\""
			i += 1
			continue
		case 'x':
			x, _ := strconv.ParseUint(source[i+2:i+4], 16, 0)
			result += fmt.Sprintf("%c", x)
			i += 3
			continue
		default:
		}
	}

	return result
}

func encodeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = fmt.Sprintf(`"%s"`, s)
	return s
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}

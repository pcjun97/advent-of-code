package aoc

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput() string {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "invalid number of arguments")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("error reading input")
	}

	return strings.TrimSuffix(string(data), "\n")
}

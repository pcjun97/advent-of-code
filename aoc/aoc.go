package aoc

import (
	"os"
	"path"
	"strings"
)

func ReadInput() string {
	dir := path.Dir(os.Args[0])
	inputFile := path.Join(dir, "input")

	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic("error reading input")
	}

	return strings.TrimSuffix(string(data), "\n")
}

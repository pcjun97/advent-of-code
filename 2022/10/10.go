package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	cpu := NewCPU(1, input)

	sum := 0
	c := []int{20, 60, 100, 140, 180, 220}
	var crt [40 * 6]bool

	for i := 1; i <= 40*6; i++ {
		for _, cycle := range c {
			if i == cycle {
				sum += i * cpu.X
			}
		}

		p := (i - 1) % 40
		if p >= cpu.X-1 && p <= cpu.X+1 {
			crt[i-1] = true
		}

		cpu.Tick()
	}

	fmt.Println(sum)
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if crt[i*40+j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

var executionCycles = map[string]int{
	"addx": 2,
	"noop": 1,
}

type CPU struct {
	X            int
	wait         int
	pc           int
	instructions []Instruction
}

func NewCPU(x int, s string) *CPU {
	instructions := make([]Instruction, 0)

	for _, line := range strings.Split(s, "\n") {
		fields := strings.Split(line, " ")
		name := fields[0]
		arg := 0

		if len(fields) > 1 {
			val, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}
			arg = val
		}

		ins := Instruction{
			Name:   name,
			Arg:    arg,
			Cycles: executionCycles[name],
		}

		instructions = append(instructions, ins)
	}

	cpu := CPU{
		X:            x,
		instructions: instructions,
		pc:           0,
		wait:         instructions[0].Cycles - 1,
	}
	return &cpu
}

func (cpu *CPU) Tick() {
	if cpu.pc >= len(cpu.instructions)-1 {
		return
	}

	if cpu.wait > 0 {
		cpu.wait--
		return
	}

	cpu.Execute()
	cpu.FetchNext()
}

func (cpu *CPU) FetchNext() {
	if cpu.pc < len(cpu.instructions)-1 {
		cpu.pc++
		cpu.wait = cpu.instructions[cpu.pc].Cycles - 1
	}
}

func (cpu *CPU) Execute() {
	ins := cpu.instructions[cpu.pc]
	switch ins.Name {
	case "addx":
		cpu.AddX(ins.Arg)
	case "noop":
		return
	}
}

func (cpu *CPU) AddX(val int) {
	cpu.X += val
}

type Instruction struct {
	Name   string
	Arg    int
	Cycles int
}

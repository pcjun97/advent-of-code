package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type InstructionType int

const (
	HLF InstructionType = iota
	TPL
	INC
	JMP
	JIE
	JIO
)

type Instruction struct {
	InstructionType
	register byte
	value    int
}

func parseInstruction(s string) Instruction {
	r := regexp.MustCompile(`(\w+) ([^,]+)(, )?([^,]*)?`)
	m := r.FindStringSubmatch(s)

	ins := Instruction{}
	switch m[1] {
	case "hlf":
		ins.InstructionType = HLF
		ins.register = m[2][0]
	case "tpl":
		ins.InstructionType = TPL
		ins.register = m[2][0]
	case "inc":
		ins.InstructionType = INC
		ins.register = m[2][0]
	case "jmp":
		ins.InstructionType = JMP
		offset, _ := strconv.Atoi(m[2])
		ins.value = offset
	case "jie":
		ins.InstructionType = JIE
		ins.register = m[2][0]
		offset, _ := strconv.Atoi(m[4])
		ins.value = offset
	case "jio":
		ins.InstructionType = JIO
		ins.register = m[2][0]
		offset, _ := strconv.Atoi(m[4])
		ins.value = offset
	}
	return ins
}

type Computer struct {
	counter      int
	registers    map[byte]uint
	instructions []Instruction
}

func NewComputer(instructions []Instruction) *Computer {
	c := Computer{
		counter:      0,
		registers:    make(map[byte]uint),
		instructions: instructions,
	}
	return &c
}

func (c *Computer) Reset() {
	c.counter = 0
	c.registers = make(map[byte]uint)
}

func (c *Computer) StartExecution() {
	for c.counter < len(c.instructions) {
		c.Execute()
	}
}

func (c *Computer) Execute() {
	ins := c.instructions[c.counter]
	r := ins.register
	n := ins.value

	switch ins.InstructionType {
	case HLF:
		v, ok := c.registers[r]
		if !ok {
			v = 0
		}
		c.registers[r] = v / 2
		c.counter += 1
	case TPL:
		v, ok := c.registers[r]
		if !ok {
			v = 0
		}
		c.registers[r] = v * 3
		c.counter += 1
	case INC:
		v, ok := c.registers[r]
		if !ok {
			v = 0
		}
		c.registers[r] = v + 1
		c.counter += 1
	case JMP:
		c.counter += n
	case JIE:
		if (c.registers[r] % 2) == 0 {
			c.counter += n
		} else {
			c.counter += 1
		}
	case JIO:
		if c.registers[r] == 1 {
			c.counter += n
		} else {
			c.counter += 1
		}
	}
}

func (c *Computer) GetRegisterValue(r byte) uint {
	return c.registers[r]
}

func (c *Computer) SetRegisterValue(r byte, v uint) {
	c.registers[r] = v
}

type Solver struct {
	computer *Computer
}

func NewSolver(input string) *Solver {
	instructions := []Instruction{}
	for _, line := range strings.Split(input, "\n") {
		instructions = append(instructions, parseInstruction(line))
	}
	c := NewComputer(instructions)
	s := Solver{c}
	return &s
}

func (s *Solver) Part1() int {
	s.computer.Reset()
	s.computer.StartExecution()
	return int(s.computer.GetRegisterValue('b'))
}

func (s *Solver) Part2() int {
	s.computer.Reset()
	s.computer.SetRegisterValue('a', 1)
	s.computer.StartExecution()
	return int(s.computer.GetRegisterValue('b'))
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}

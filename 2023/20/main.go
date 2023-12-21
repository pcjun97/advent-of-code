package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type PulseType int

const (
	LowPulse PulseType = iota
	HighPulse
)

type Pulse struct {
	t        PulseType
	from, to string
}

type Module interface {
	ID() string
	Send(p Pulse)
	AddInput(m Module)
	AddOutput(m Module)
	ResolvePulses() []Pulse
	Reset()
	HasOutput(id string) bool
}

type BaseModule struct {
	id      string
	pulses  []Pulse
	outputs []Module
}

func NewBaseModule(id string) *BaseModule {
	m := BaseModule{id, nil, nil}
	return &m
}

func (m *BaseModule) ID() string {
	return m.id
}

func (m *BaseModule) HasOutput(id string) bool {
	for _, output := range m.outputs {
		if output.ID() == id {
			return true
		}
	}
	return false
}

func (m *BaseModule) Send(p Pulse) {
	m.pulses = append(m.pulses, p)
}

func (m *BaseModule) AddInput(input Module) {}

func (m *BaseModule) AddOutput(output Module) {
	m.outputs = append(m.outputs, output)
}

func (m *BaseModule) ResolvePulses() []Pulse {
	m.pulses = nil
	return nil
}

func (m *BaseModule) Reset() {
	m.pulses = nil
}

type FlipFlopModule struct {
	BaseModule
	out PulseType
}

func NewFlipFlopModule(id string) *FlipFlopModule {
	base := BaseModule{id, nil, nil}
	m := FlipFlopModule{base, LowPulse}
	return &m
}

func (m *FlipFlopModule) ResolvePulses() []Pulse {
	pulses := []Pulse{}

	for _, p := range m.pulses {
		if p.t == HighPulse {
			continue
		}

		switch m.out {
		case HighPulse:
			m.out = LowPulse
		case LowPulse:
			m.out = HighPulse
		}

		for _, output := range m.outputs {
			p := Pulse{m.out, m.ID(), output.ID()}
			pulses = append(pulses, p)
		}
	}

	m.pulses = nil

	return pulses
}

func (m *FlipFlopModule) Reset() {
	m.BaseModule.Reset()
	m.out = LowPulse
}

type ConjunctionModule struct {
	BaseModule
	mem map[string]PulseType
}

func NewConjunctionModule(id string) *ConjunctionModule {
	base := BaseModule{id, nil, nil}
	m := ConjunctionModule{base, make(map[string]PulseType)}
	return &m
}

func (m *ConjunctionModule) AddInput(input Module) {
	m.mem[input.ID()] = LowPulse
}

func (m *ConjunctionModule) ResolvePulses() []Pulse {
	pulses := []Pulse{}

	for _, pulse := range m.pulses {
		m.mem[pulse.from] = pulse.t

		out := LowPulse
		for _, t := range m.mem {
			if t == LowPulse {
				out = HighPulse
				break
			}
		}

		for _, output := range m.outputs {
			pulse := Pulse{out, m.ID(), output.ID()}
			pulses = append(pulses, pulse)
		}
	}
	m.pulses = nil

	return pulses
}

func (m *ConjunctionModule) Reset() {
	m.BaseModule.Reset()
	for mod := range m.mem {
		m.mem[mod] = LowPulse
	}
}

type BroadcastModule struct {
	BaseModule
}

func NewBroadcastModule() *BroadcastModule {
	base := BaseModule{"broadcaster", nil, nil}
	m := BroadcastModule{base}
	return &m
}

func (m *BroadcastModule) ResolvePulses() []Pulse {
	pulses := []Pulse{}

	for _, pulse := range m.pulses {
		for _, output := range m.outputs {
			pulses = append(pulses, Pulse{pulse.t, m.ID(), output.ID()})
		}
	}
	m.pulses = nil

	return pulses
}

func parseModules(s string) map[string]Module {
	modules := make(map[string]Module)

	r := regexp.MustCompile(`(%|&)?(.+) -> (.+)`)
	mm := [][]string{}
	for _, line := range strings.Split(s, "\n") {
		mm = append(mm, r.FindStringSubmatch(line))
	}

	for _, m := range mm {
		id := m[2]
		switch {
		case m[1] == "%":
			modules[id] = NewFlipFlopModule(id)
		case m[1] == "&":
			modules[id] = NewConjunctionModule(id)
		case id == "broadcaster":
			modules[id] = NewBroadcastModule()
		default:
			modules[id] = NewBaseModule(id)
		}
	}

	for _, m := range mm {
		module := modules[m[2]]
		for _, id := range strings.Split(m[3], ", ") {
			output, ok := modules[id]
			if !ok {
				output = NewBaseModule(id)
				modules[id] = output
			}
			module.AddOutput(output)
			output.AddInput(module)
		}
	}

	return modules
}

type Solver struct {
	modules map[string]Module
}

func NewSolver(input string) *Solver {
	s := Solver{parseModules(input)}
	return &s
}

func (s *Solver) Part1() int {
	low, high := 0, 0

	for i := 0; i < 1000; i++ {
		button := Pulse{LowPulse, "", "broadcaster"}
		pulses := []Pulse{button}

		for len(pulses) > 0 {
			next := []Pulse{}

			for _, pulse := range pulses {
				switch pulse.t {
				case LowPulse:
					low += 1
				case HighPulse:
					high += 1
				}

				s.modules[pulse.to].Send(pulse)
			}
			for _, module := range s.modules {
				next = append(next, module.ResolvePulses()...)
			}

			pulses = next
		}
	}

	for _, mod := range s.modules {
		mod.Reset()
	}

	return low * high
}

func (s *Solver) Part2() int {
	rx := s.modules["rx"]

	var prev Module
	for _, module := range s.modules {
		if module.HasOutput(rx.ID()) {
			prev = module
		}
	}

	inputs := []Module{}
	for _, module := range s.modules {
		if module.HasOutput(prev.ID()) {
			inputs = append(inputs, module)
		}
	}

	cache := make(map[string]int)

	i := 1
	for {
		button := Pulse{LowPulse, "", "broadcaster"}
		pulses := []Pulse{button}

		for len(pulses) > 0 {
			next := []Pulse{}

			for _, pulse := range pulses {
				if pulse.to == prev.ID() && pulse.t == HighPulse {
					if _, ok := cache[pulse.from]; !ok {
						cache[pulse.from] = i
						if len(cache) == len(inputs) {
							result := 1
							for _, v := range cache {
								result = aoc.LCM(result, v)
							}
							return result
						}
					}
				}

				s.modules[pulse.to].Send(pulse)
			}
			for _, module := range s.modules {
				next = append(next, module.ResolvePulses()...)
			}

			pulses = next
		}

		i += 1
	}
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}

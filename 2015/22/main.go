package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Character struct {
	hp, damage int
}

type Solver struct {
	boss Character
}

func parseBoss(s string) Character {
	r := regexp.MustCompile(`Hit Points: (.*)\nDamage: (.*)$`)
	m := r.FindStringSubmatch(s)
	hp, _ := strconv.Atoi(m[1])
	damage, _ := strconv.Atoi(m[2])
	boss := Character{hp, damage}
	return boss
}

func NewSolver(input string) *Solver {
	s := Solver{parseBoss(input)}
	return &s
}

func (s *Solver) Part1(hp, mana int) int {
	self := Character{hp, 0}
	initial := State{
		self: self,
		boss: s.boss,
		mana: mana,
	}

	states := []State{initial}
	cache := map[State]struct{}{}
	for len(states) > 0 {
		min := 0
		for i, state := range states {
			if state.spent < states[min].spent {
				min = i
			}
		}
		state := states[min]
		states = append(states[:min], states[min+1:]...)
		if _, ok := cache[state]; ok {
			continue
		}
		cache[state] = struct{}{}

		if state.BossIsDead() {
			return state.spent
		}
		states = append(states, state.NextStates()...)
	}

	return 0
}

func (s *Solver) Part2(hp, mana int) int {
	self := Character{hp, 0}
	initial := State{
		self: self,
		boss: s.boss,
		mana: mana,
	}

	states := []State{initial}
	cache := map[State]struct{}{}
	for len(states) > 0 {
		min := 0
		for i, state := range states {
			if state.spent < states[min].spent {
				min = i
			}
		}
		state := states[min]
		states = append(states[:min], states[min+1:]...)
		if _, ok := cache[state]; ok {
			continue
		}
		cache[state] = struct{}{}

		if state.BossIsDead() {
			return state.spent
		}
		if state.turn == 0 {
			state.self.hp -= 1
		}
		states = append(states, state.NextStates()...)
	}

	return 0
}

type State struct {
	turn     int
	self     Character
	boss     Character
	mana     int
	spent    int
	shield   int
	poison   int
	recharge int
}

func (s State) IsFinal() bool {
	return s.self.hp <= 0 || s.boss.hp <= 0
}

func (s State) BossIsDead() bool {
	return s.boss.hp <= 0
}

func (s State) NextStates() []State {
	if s.IsFinal() {
		return nil
	}

	next := s
	next.turn = (next.turn + 1) % 2

	if next.poison > 0 {
		next.boss.hp -= 3
		next.poison -= 1
	}

	if next.recharge > 0 {
		next.mana += 101
		next.recharge -= 1
	}

	shielded := false
	if next.shield > 0 {
		shielded = true
		next.shield -= 1
	}

	if next.turn == 0 {
		dmg := next.boss.damage
		if shielded {
			dmg -= 7
		}
		if dmg <= 0 {
			dmg = 1
		}
		next.self.hp -= dmg
		return []State{next}
	}

	states := []State{}
	if next.mana >= 53 {
		n := next
		n.mana -= 53
		n.spent += 53
		n.boss.hp -= 4
		states = append(states, n)
	}

	if next.mana >= 73 {
		n := next
		n.mana -= 73
		n.spent += 73
		n.boss.hp -= 2
		n.self.hp += 2
		states = append(states, n)
	}

	if next.mana >= 113 && next.shield <= 0 {
		n := next
		n.mana -= 113
		n.spent += 113
		n.shield = 6
		states = append(states, n)
	}

	if next.mana >= 173 && next.poison <= 0 {
		n := next
		n.mana -= 173
		n.spent += 173
		n.poison = 6
		states = append(states, n)
	}

	if next.mana >= 229 && next.recharge <= 0 {
		n := next
		n.mana -= 229
		n.spent += 229
		n.recharge = 5
		states = append(states, n)
	}

	return states
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1(50, 500))
	fmt.Println(s.Part2(50, 500))
}

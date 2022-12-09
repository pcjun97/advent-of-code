package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	tg := NewTreeGrid(input)

	fmt.Println(len(tg.VisibleTrees()))
	fmt.Println(tg.HighestScoreTree().ScenicScore())
}

type TreeGrid [][]*Tree

func NewTreeGrid(s string) TreeGrid {
	lines := strings.Split(s, "\n")

	row := len(lines)
	col := len(lines[0])

	tg := make(TreeGrid, len(lines))

	for i, line := range lines {
		tg[i] = make([]*Tree, len(line))
		for j, h := range line {
			height := int(h - '0')
			tg[i][j] = NewTree(j, i, height)
		}
	}

	var max int
	if col > row {
		max = col
	} else {
		max = row
	}

	var tr [4]Tracker
	var t [4]*Tree

	for i := 0; i < max; i++ {
		for k := 0; k < 4; k++ {
			tr[k] = NewTracker()
		}

		for j := 0; j < max; j++ {
			for k := 0; k < 4; k++ {
				t[k] = nil
			}

			if i < row && j < col {
				t[0] = tg[i][j]
			}

			if i < row && max-j-1 < col {
				t[1] = tg[i][max-j-1]
			}

			if j < row && i < col {
				t[2] = tg[j][i]
			}

			if max-j-1 < row && i < col {
				t[3] = tg[max-j-1][i]
			}

			for k := 0; k < 4; k++ {
				if t[k] == nil {
					continue
				}
				bt := tr[k].BlockingTree(t[k])
				if bt == nil {
					t[k].Visibility[k] = true
					t[k].Score[k] = j
				} else {
					t[k].Visibility[k] = false
					t[k].Score[k] = t[k].DistanceX(bt) + t[k].DistanceY(bt)
				}
				tr[k].Add(t[k])
			}
		}
	}

	return tg
}

func (tg TreeGrid) VisibleTrees() []*Tree {
	v := []*Tree{}
	for i := 0; i < len(tg); i++ {
		for j := 0; j < len(tg[i]); j++ {
			if tg[i][j].Visible() {
				v = append(v, tg[i][j])
			}
		}
	}
	return v
}

func (tg TreeGrid) HighestScoreTree() *Tree {
	t := (*Tree)(nil)
	for i := 0; i < len(tg); i++ {
		for j := 0; j < len(tg[i]); j++ {
			if t == nil || tg[i][j].ScenicScore() > t.ScenicScore() {
				t = tg[i][j]
			}
		}
	}
	return t
}

type Tree struct {
	X          int
	Y          int
	Height     int
	Visibility [4]bool
	Score      [4]int
}

func NewTree(x, y, height int) *Tree {
	t := Tree{
		X:      x,
		Y:      y,
		Height: height,
	}

	return &t
}

func (t *Tree) Visible() bool {
	for _, v := range t.Visibility {
		if v {
			return true
		}
	}

	return false
}

func (t *Tree) ScenicScore() int {
	score := 1

	for _, s := range t.Score {
		score *= s
	}

	return score
}

func (t *Tree) DistanceX(other *Tree) int {
	d := t.X - other.X
	if d < 0 {
		return -d
	}
	return d
}

func (t *Tree) DistanceY(other *Tree) int {
	d := t.Y - other.Y
	if d < 0 {
		return -d
	}
	return d
}

type Tracker map[*Tree]bool

func NewTracker() Tracker {
	tr := make(Tracker)
	return tr
}

func (tr Tracker) Add(t *Tree) {
	for o := range tr {
		if o.Height <= t.Height {
			delete(tr, o)
		}
	}
	tr[t] = true
}

func (tr Tracker) BlockingTree(t *Tree) *Tree {
	bt := (*Tree)(nil)
	for o := range tr {
		if o.Height >= t.Height {
			if bt == nil {
				bt = o
			} else if o.Height < bt.Height {
				bt = o
			}
		}
	}
	return bt
}

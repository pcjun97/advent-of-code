package main

import (
	"fmt"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type FileBlock struct {
	ID         int
	Length     int
	Prev, Next *FileBlock
}

func NewFileBlock() *FileBlock {
	block := FileBlock{
		ID:     -1,
		Length: 0,
	}

	return &block
}

type DiskMap struct {
	Head, Tail *FileBlock
}

func ParseDiskMap(s string) DiskMap {
	cur := NewFileBlock()
	head := cur
	tail := cur

	id := 0
	free := false
	for _, r := range s {
		tail = cur

		length := int(r - '0')
		cur.Length = length

		if !free {
			cur.ID = id
			id += 1
		}

		next := NewFileBlock()
		next.Prev = cur
		cur.Next = next

		cur = next
		free = !free
	}

	tail.Next.Prev = nil
	tail.Next = nil

	dm := DiskMap{head, tail}
	dm.Compact()

	return dm
}

func (d *DiskMap) Clone() DiskMap {
	dm := DiskMap{}
	if d.Head == nil {
		return dm
	}

	cur := NewFileBlock()
	head := cur
	tail := cur

	from := d.Head
	for from != nil {
		tail = cur

		cur.ID = from.ID
		cur.Length = from.Length

		next := NewFileBlock()
		next.Prev = cur
		cur.Next = next

		cur = next
		from = from.Next
	}

	tail.Next.Prev = nil
	tail.Next = nil

	dm.Head = head
	dm.Tail = tail
	return dm
}

func (d *DiskMap) Defrag() DiskMap {
	defrag := d.Clone()
	if defrag.Head == nil || defrag.Tail == nil {
		return defrag
	}

	// add empty block at the end
	if defrag.Tail.ID >= 0 {
		block := NewFileBlock()

		defrag.Tail.Next = block
		block.Prev = defrag.Tail

		defrag.Tail = block
	}

	free := defrag.Tail

	// actual defrag
	cur := defrag.Head
	for cur != defrag.Tail {
		if cur.ID >= 0 {
			cur = cur.Next
			continue
		}

		last := free.Prev

		switch {
		case cur.Length == last.Length:
			cur.ID = last.ID

			free.Length += last.Length
			free.Prev = last.Prev
			free.Prev.Next = free

		case cur.Length < last.Length:
			cur.ID = last.ID

			last.Length -= cur.Length
			free.Length += cur.Length

		case cur.Length > last.Length:
			free.Length += last.Length
			free.Prev = last.Prev
			free.Prev.Next = free

			last.Prev = cur.Prev
			last.Next = cur

			last.Prev.Next = last
			cur.Prev = last

			cur.Length -= last.Length
			cur = last
		}

		for free.Prev.ID < 0 {
			free.Length += free.Prev.Length
			free.Prev = free.Prev.Prev
			free.Prev.Next = free
		}

		cur = cur.Next
	}

	return defrag
}

func (d *DiskMap) DefragWhole() DiskMap {
	defrag := d.Clone()
	if defrag.Head == nil || defrag.Tail == nil {
		return defrag
	}

	t := make(map[int]struct{})

	cur := defrag.Tail
	for cur != nil {
		if cur.ID < 0 {
			cur = cur.Prev
			continue
		}

		if _, ok := t[cur.ID]; ok {
			cur = cur.Prev
			continue
		}
		t[cur.ID] = struct{}{}

		free := defrag.Head
		for free != nil && free != cur && (free.ID >= 0 || free.Length < cur.Length) {
			free = free.Next
		}

		if free == nil || free == cur {
			cur = cur.Prev
			continue
		}

		moved := NewFileBlock()
		moved.ID = cur.ID
		moved.Length = cur.Length
		moved.Prev = free.Prev
		moved.Next = free

		if moved.Prev != nil {
			moved.Prev.Next = moved
		}

		free.Prev = moved
		free.Length -= moved.Length

		cur.ID = -1
		cur = cur.Prev

		defrag.Compact()
	}

	return defrag
}

func (d *DiskMap) Compact() {
	if d.Head == nil {
		return
	}

	for d.Head.Length <= 0 {
		d.Head = d.Head.Next
	}
	d.Head.Prev = nil

	cur := d.Head.Next
	for cur != nil {
		if cur.Length <= 0 {
			cur.Prev.Next = cur.Next
			if cur.Next != nil {
				cur.Next.Prev = cur.Prev
			}
			cur = cur.Next
			continue
		}

		if cur.Prev.ID != cur.ID {
			cur = cur.Next
			continue
		}

		cur.Prev.Length += cur.Length
		cur.Prev.Next = cur.Next

		if cur.Next != nil {
			cur.Next.Prev = cur.Prev
		}

		d.Tail = cur
		cur = cur.Next
	}
}

func (d DiskMap) Checksum() int {
	sum := 0

	i := 0
	cur := d.Head

	for cur != nil {
		for j := 0; j < cur.Length; j++ {
			if cur.ID >= 0 {
				sum += cur.ID * (i)
			}

			i++
		}
		cur = cur.Next
	}

	return sum
}

func (d DiskMap) ToString() string {
	s := ""

	cur := d.Head
	for cur != nil {
		for i := 0; i < cur.Length; i++ {
			if cur.ID < 0 {
				s += fmt.Sprint(".")
			} else {
				s += fmt.Sprint(cur.ID)
			}
		}
		cur = cur.Next
	}

	return s
}

type Solver struct {
	Diskmap DiskMap
}

func NewSolver(input string) *Solver {
	s := Solver{ParseDiskMap(input)}
	return &s
}

func (s *Solver) Part1() int {
	defrag := s.Diskmap.Defrag()
	return defrag.Checksum()
}

func (s *Solver) Part2() int {
	defrag := s.Diskmap.DefragWhole()
	return defrag.Checksum()
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}

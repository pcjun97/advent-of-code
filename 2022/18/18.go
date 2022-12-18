package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	d := NewDroplet(input)
	fmt.Println(d.TotalSurfaceArea())
	fmt.Println(d.ExternalSurfaceAres())
}

type Droplet map[int]map[int]map[int]bool

func NewDroplet(s string) Droplet {
	d := make(Droplet)
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Split(line, ",")

		x, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		z, err := strconv.Atoi(fields[2])
		if err != nil {
			panic(err)
		}

		d.AddCube(x, y, z)
	}
	return d
}

func (d Droplet) AddCube(x, y, z int) {
	if _, ok := d[x]; !ok {
		d[x] = make(map[int]map[int]bool)
	}
	if _, ok := d[x][y]; !ok {
		d[x][y] = make(map[int]bool)
	}
	d[x][y][z] = true
}

func (d Droplet) HasCube(x, y, z int) bool {
	if _, ok := d[x]; !ok {
		return false
	}
	if _, ok := d[x][y]; !ok {
		return false
	}
	return d[x][y][z]
}

func (d Droplet) TotalSurfaceArea() int {
	iteration := [][3]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}

	sum := 0
	for x := range d {
		for y := range d[x] {
			for z := range d[x][y] {
				for _, n := range iteration {
					if !d.HasCube(x+n[0], y+n[1], z+n[2]) {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func (d Droplet) ExternalSurfaceAres() int {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	minZ := math.MaxInt
	maxZ := math.MinInt

	for x := range d {
		if x >= maxX {
			maxX = x + 1
		}
		if x <= minX {
			minX = x - 1
		}
		for y := range d[x] {
			if y >= maxY {
				maxY = y + 1
			}
			if y <= minY {
				minY = y - 1
			}
			for z := range d[x][y] {
				if z >= maxZ {
					maxZ = z + 1
				}
				if z <= minZ {
					minZ = z - 1
				}
			}
		}
	}

	iteration := [][3]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}

	sum := 0
	visited := make(map[[3]int]bool)
	tovisit := make(map[[3]int]bool)
	tovisit[[3]int{minX, minY, minZ}] = true
	for len(tovisit) > 0 {
		for v := range tovisit {
			x := v[0]
			y := v[1]
			z := v[2]

			for _, n := range iteration {
				nx := x + n[0]
				ny := y + n[1]
				nz := z + n[2]
				if nx < minX || nx > maxX || ny < minY || ny > maxY || nz < minZ || nz > maxZ {
					continue
				}

				if d.HasCube(nx, ny, nz) {
					sum++
				} else if key := [3]int{nx, ny, nz}; !visited[key] {
					tovisit[key] = true
				}
			}

			delete(tovisit, v)
			visited[[3]int{x, y, z}] = true
		}
	}
	return sum
}

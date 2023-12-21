package aoc

type Range struct {
	Min, Max int
}

func NewRange(min, max int) Range {
	r := Range{min, max}
	return r
}

func (self Range) Overlap(r Range) bool {
	if self.Max < r.Min || r.Max < self.Min {
		return false
	}
	return true
}

func (self Range) Dedup(r Range) []Range {
	if self.Min == r.Min && self.Max == r.Max {
		return []Range{self}
	}

	a, b := self, r
	if b.Min < a.Min {
		a, b = b, a
	}
	if a.Min == b.Min && b.Max < a.Max {
		a, b = b, a
	}

	if !a.Overlap(b) {
		return []Range{a, b}
	}

	if a.Min == b.Min {
		return []Range{a, {a.Max + 1, b.Max}}
	}

	if a.Max == b.Max {
		return []Range{{a.Min, b.Min - 1}, b}
	}

	if b.Min > a.Min && b.Max < a.Max {
		return []Range{{a.Min, b.Min - 1}, b, {b.Max + 1, a.Max}}
	}

	return []Range{{a.Min, b.Min - 1}, {b.Min, a.Max}, {a.Max + 1, b.Max}}
}

func AddDedupRanges(ranges []Range, r Range) []Range {
	result := []Range{}

	for _, rr := range ranges {
		dedup := r.Dedup(rr)
		r = dedup[len(dedup)-1]
		result = append(result, dedup[:len(dedup)-1]...)
	}
	result = append(result, r)

	return result
}

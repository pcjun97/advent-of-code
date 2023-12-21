package aoc

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func Extrapolate(list []int) int {
	if len(list) == 0 {
		return 0
	}

	if len(list) == 1 {
		return list[0]
	}

	diff := []int{}
	for i := 1; i < len(list); i++ {
		diff = append(diff, list[i]-list[i-1])
	}

	zero := true
	for _, d := range diff {
		if d != 0 {
			zero = false
			break
		}
	}

	if zero {
		return list[0]
	}

	next := Extrapolate(diff)

	return list[len(list)-1] + next
}

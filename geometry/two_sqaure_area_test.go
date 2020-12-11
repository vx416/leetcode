package geometry

type Sqaure struct {
	minX, maxX, minY, maxY int
}

func (s Sqaure) xlen() int {
	return s.maxX - s.minX
}

func (s Sqaure) ylen() int {
	return s.maxY - s.minY
}

func (s Sqaure) Area() int {
	return s.xlen() * s.ylen()
}

func (s Sqaure) Overlap(other Sqaure) *Sqaure {
	if s.maxX > other.maxX || s.maxY > other.maxY {
		return nil
	}
	if s.minX < other.minX || s.minY < other.minY {
		return nil
	}

	maxX := min(s.maxX, other.maxX)
	minX := max(s.minX, other.minX)
	maxY := min(s.maxY, other.maxY)
	minY := max(s.minY, other.minY)
	return &Sqaure{maxX: maxX, minX: minX, maxY: maxY, minY: minY}
}

func max(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > res {
			res = a[i]
		}
	}
	return res
}

func min(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}

func twoSqaureAres(a, b [][2]int) int {
	return 0
}

package pattern

// Strategy interface (find pivot`s slice element for quick sort)
type FindPivotStrategy interface {
	FindPivot(l, h int, a []int) int
}

// Concrete implementations
type MiddlePivot struct{}

func (m *MiddlePivot) FindPivot(l, h int, a []int) int {
	return a[(l+h)/2]
}

type LomutoPivot struct{}

func (avg *LomutoPivot) FindPivot(l, h int, a []int) int {
	m := (l + h) / 2
	if a[m] < a[l] {
		a[m], a[l] = a[l], a[m]
	}
	if a[h] < a[l] {
		a[h], a[l] = a[l], a[h]
	}
	if a[h] < a[m] {
		a[h], a[m] = a[m], a[h]
	}
	return a[m]
}

// Some context for quick sorting
type SortContext struct {
	pivotAlgorithm FindPivotStrategy
}

func (s *SortContext) SwitchStrategy(st FindPivotStrategy) {
	s.pivotAlgorithm = st
}

func (s *SortContext) GetListPivot(l, h int, a []int) int {
	return s.pivotAlgorithm.FindPivot(l, h, a)
}

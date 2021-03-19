package quicksort

func QuickSort(xi []int) []int {
	if len(xi) < 2 {
		return xi
	}
	pivot := xi[0]
	var less, greater []int

	for _, v := range xi[1:] {
		if v <= pivot {
			less = append(less, v)
		} else if v > pivot {
			greater = append(greater, v)
		}
	}
	var sorted []int
	sorted = append(sorted, QuickSort(less)...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, QuickSort(greater)...)

	return sorted
}

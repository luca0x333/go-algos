package binary_search

import "testing"

func BenchmarkBinarySearchSwitch(b *testing.B) {
	xi := []int{1, 2, 4, 3, 2, 45, 65, 32, 143, 67, 32, 23}
	for i := 0; i < b.N; i++ {
		BinarySearchSwitch(xi, 67)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	xi := []int{1, 2, 4, 3, 2, 45, 65, 32, 143, 67, 32, 23}
	for i := 0; i < b.N; i++ {
		BinarySearch(xi, 67)
	}
}

func BenchmarkSimpleSearch(b *testing.B) {
	xi := []int{1, 2, 4, 3, 2, 45, 65, 32, 143, 67, 32, 23}
	for i := 0; i < b.N; i++ {
		SimpleSearch(xi, 67)
	}
}

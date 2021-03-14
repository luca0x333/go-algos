package binary_search

import "sort"

func BinarySearchSwitch(xi []int, n int) int {
	// return -1 if n is not present
	r := -1

	sort.Ints(xi)

	low := 0
	high := len(xi) - 1

	switch {
	case len(xi) == 0:
		return -1
	case xi[0] == n:
		return n
	default:
		for low <= high {
			mid := (low + high) / 2
			if xi[mid] == n {
				return n
			} else if xi[mid] < n {
				low = mid + 1
			} else if xi[mid] > n {
				high = mid - 1
			}
		}
		return r
	}
}

func BinarySearch(xi []int, n int) bool {
	sort.Ints(xi)
	low := 0
	high := len(xi) - 1
	for low <= high {
		mid := (low + high) / 2
		if xi[mid] == n {
			return true
		} else if xi[mid] < n {
			low = mid + 1
		} else if xi[mid] > n {
			high = mid - 1
		}
	}
	return false
}

func SimpleSearch(xi []int, n int) bool {
	for _, v := range xi {
		if v == n {
			return true
		}
	}
	return false
}

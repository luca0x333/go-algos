package main

import (
	"fmt"
	"github.com/luca0x333/go-algos/binary_search"
)

func main() {
	xi := []int{1, 2, 4, 3, 2, 45, 65, 32, 143, 67, 32, 23}

	fmt.Println(binary_search.BinarySearchSwitch(xi, 32))
	fmt.Println(binary_search.BinarySearch(xi, 32))
	fmt.Println(binary_search.SimpleSearch(xi, 32))
}

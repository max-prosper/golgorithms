package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1000, 10, 9, 8, 7, 6, 4, 5, 3, 2, 1, 100}
	mergeSort(&nums, 0, len(nums)-1)

	fmt.Printf("SORTED: %v \n", nums)
}

func merge(nums *[]int, s, m, e int) *[]int {

	n1 := m - s + 1
	n2 := e - m

	left := make([]int, n1)
	right := make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = (*nums)[s+i]
	}

	for j := 0; j < n2; j++ {
		right[j] = (*nums)[m+j+1]
	}

	left = append(left, math.MaxUint32)
	right = append(right, math.MaxUint32)

	i := 0
	j := 0

	for k := s; k <= e; k++ {
		if left[i] <= right[j] {
			(*nums)[k] = left[i]
			i++
		} else {
			(*nums)[k] = right[j]
			j++
		}
	}
	return nums
}

func mergeSort(nums *[]int, s, e int) *[]int {
	if s < e {
		m := int((s + e) / 2)
		mergeSort(nums, s, m)
		mergeSort(nums, m+1, e)
		merge(nums, s, m, e)

		return nums
	}
	return nums
}

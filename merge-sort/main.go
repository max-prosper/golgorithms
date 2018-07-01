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

// this helper function creates two slices based on the length and values of the first and second halves of an original slice
// then values of these two slice compared one by one
// on each iteration, smaller value copied to the original slice
func merge(nums *[]int, s, m, e int) *[]int {

	n1 := m - s + 1 // calculate the length of the left slice and right slices
	n2 := e - m     // 's' - start, 'm' - middle, 'e' - end

	left, right := make([]int, n1), make([]int, n2) // make two slices of length n1 and n2

	for i := 0; i < n1; i++ {
		left[i] = (*nums)[s+i] // fill first slice with values of the first half of 'nums' slice
	}

	for j := 0; j < n2; j++ {
		right[j] = (*nums)[m+j+1] // fill second slice with values of the second half of 'nums' slice
	}

	left = append(left, math.MaxUint32)   // add a 'signal' value (max possible uint32) to the end of both slices
	right = append(right, math.MaxUint32) // any integer can't be bigger than this value

	i, j := 0, 0 // initialize variable to work with indecies

	for k := s; k <= e; k++ {
		if left[i] <= right[j] { // compare first values of left and rigth slices
			(*nums)[k] = left[i] // and write a smaller one to 'sum' slice at the index of iteration
			i++                  // increase the index for the slice smaller value was copied from
		} else {
			(*nums)[k] = right[j]
			j++
		}
	}
	return nums
}

// this function uses the helper 'merge' function in a recursion calls
func mergeSort(nums *[]int, s, e int) *[]int {
	if s < e { // if 's' is no smaller than 'e', left and right slices lenghts are 1 and they are sorted
		m := int((s + e) / 2)   // define the 'meiddle' point by which the original slice will be divided
		mergeSort(nums, s, m)   // call 'mergeSort' on the first half recursively
		mergeSort(nums, m+1, e) // call 'mergeSort' on the second half recursively
		merge(nums, s, m, e)    // use 'merge' to sort halved slice and write new soreted values to 'sum' on each recursive call and

		return nums
	}
	return nums
}

package main

import "fmt"

func insertionSort(nums []int) {
	for j := 1; j < len(nums); j++ {
		// 'key' is the element to insert
		key := nums[j]
		// 'i' is the index of the element to compare 'key' with,
		// initially it's the first element of the input slice.
		i := j - 1
		// So, technically we devide our input slice into two:
		// 1st one considered ordered because it contains
		// only one element (the first one from the input).
		// 2nd is an unsored slice which contains the rest of the input.
		// We wll be inserting every element of 2nd slice
		// in to the end of the 1nd one

		// Compare each element of sorted
		// slice to inserted element
		for i >= 0 && nums[i] > key {
			// if compared element is bigger than the 'key' ,
			// move former by one position to the right
			nums[i+1] = nums[i]
			// and then decrease 'i' by 1 to compare
			// next element to the 'key'
			i--
		}
		// put 'key' to the former position of last bigger element
		nums[i+1] = key
	}

	fmt.Printf("Sorted: %v \n", nums)
}

func main() {
	nums := []int{1000, 10, 9, 8, 7, 6, 4, 5, 3, 2, 1, 100}
	insertionSort(nums)
}

package main

import "fmt"

func swap(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func partition(s []int) int {
	p := s[len(s)-1]
	i := -1

	for j := 0; j < len(s)-1; j++ {
		if s[j] < p {
			i++
			swap(s, i, j)
		}
	}
	swap(s, i+1, len(s)-1)
	return i + 1
}

func quickSort(s []int) {
	if len(s) > 1 {
		i := partition(s)
		quickSort(s[:i])
		quickSort(s[i+1 : len(s)-1])
	}

}

func main() {
	s := []int{1000, 10, 9, 8, 7, 6, 4, 5, 3, 2, 1, 100}
	quickSort(s)
	fmt.Printf("\nSORTED: %v\n\n", s)
}

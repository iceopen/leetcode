package _976_largest_perimeter_triangle

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}

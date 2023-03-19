package minnum

import "fmt"

// [1, 3, 6, 4, 1, 2]
// [1, 2, 3]
// [-1, -3]
func Solution(A []int) int {
	n := len(A)
	fmt.Println("n :", n)
	seen := make(map[int]bool)

	// Iterate through the array and add all positive integers to the hash set.
	for i := 0; i < n; i++ {
		if A[i] > 0 {
			seen[A[i]] = true
		}
	}

	fmt.Println("seen :", seen)
	fmt.Println("input :", A)

	// Iterate from 1 to the maximum possible value in the array, and return the first positive integer
	// that does not appear in the hash set.
	for i := 1; i <= n; i++ {
		if !seen[i] {
			return i
		}
	}

	// If all positive integers appear in the array, return the maximum value in the array plus one.
	return n + 1
}

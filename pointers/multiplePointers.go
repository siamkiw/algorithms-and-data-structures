package pointers

import (
	"errors"
	"fmt"
)

func SumZero(arr []int) ([]int, error) {
	left := 0
	right := len(arr) - 1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == 0 {
			return []int{arr[left], arr[right]}, nil
		} else if sum > 0 {
			right--
		} else {
			left++
		}
	}

	return []int{}, nil
}

// [1, 1, 3, 4, 5, 5, 6, 7]
func CounterUniqueValues(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	i := 0

	for j := 0; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
		fmt.Println("arr :", arr)
	}

	return i + 1
}

// []int{2,6,9,2,1,8,5,6,3} , 3
func MaxSumArray(arr []int, num int) (int, error) {
	maxSum := 0
	tempSum := 0
	if num > len(arr) {
		return 0, errors.New("num must greater then length of array")
	}

	for i := 0; i < num; i++ {
		maxSum += arr[i]
	}

	tempSum = maxSum

	for i := num; i < len(arr); i++ {
		tempSum = tempSum - arr[i-num] + arr[i]
		if maxSum < tempSum {
			maxSum = tempSum
		}
	}

	return maxSum, nil

}

func twoSum(nums []int, target int) []int {
	// i := 0
	// for j := 1; j < len(nums); j++ {
	//     if nums[i] + nums[j] == target {
	//         return  []int{i, j}
	//     }
	//     i++
	// }
	// return []int{}
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{0, 0}
}

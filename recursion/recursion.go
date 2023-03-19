package recursion

import "fmt"

func CountDown(num int) {
	if num <= 0 {
		fmt.Println("all done")
		return
	}
	fmt.Println(num)
	num--
	CountDown(num)
}

func SumRange(num int) int {
	if num == 1 {
		return num
	}
	return num + SumRange(num-1)
}

func Factorial(num int) int {
	if num == 1 {
		return num
	}
	return num * Factorial(num-1)
}

func CollectOddValues(arr []int) []int {
	result := []int{}

	var helper func(nums []int)
	helper = func(nums []int) {
		if len(nums) == 0 {
			return
		}

		if nums[0]%2 != 0 {
			result = append(result, nums[0])
		}

		arr = arr[1:]
		helper(arr)
	}

	helper(arr)
	return result
}

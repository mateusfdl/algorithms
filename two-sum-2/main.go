package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 9))
}

func twoSum(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1

	for start < end {
		result := nums[start] + nums[end]

		if result > target {
			end--
		} else if result < target {
			start++
		} else {
			return []int{start, end}
		}
	}

	return []int{-1, -1}
}

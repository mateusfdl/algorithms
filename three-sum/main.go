package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		start := i + 1
		end := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for start < end {
			b := nums[start]
			c := nums[end]

			if a+b+c == 0 {
				result = append(result, []int{a, b, c})

				start++
				end--
			} else if a+b+c > 0 {
				end--
			} else {
				start++
			}
		}
	}

	return result
}

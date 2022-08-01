package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {

	count := 1
	last := nums[0]
	for _, curr := range nums[1:] {
		if curr > last {
			nums[count] = curr
			last = curr
			count++
		}

	}
	return count
}

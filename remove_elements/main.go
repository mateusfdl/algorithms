package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	fmt.Println(removeElement(nums, target))
}

func removeElement(nums []int, val int) int {
	var k int

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k += 1
		}
	}

	return k
}

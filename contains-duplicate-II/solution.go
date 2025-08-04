package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 1, 2, 3}
	k := 3

	fmt.Println(containsNearbyDuplicate(i, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, num := range nums {
		if prevIndex, found := seen[num]; found {
			if i-prevIndex <= k {
				return true
			}
		}
		seen[num] = i
	}

	return false
}

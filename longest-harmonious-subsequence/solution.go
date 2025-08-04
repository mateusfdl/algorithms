package main

import "fmt"

func main() {
	s := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(s))

}
func findLHS(nums []int) int {
	freq := make(map[int]int)
	maxLength := 0

	for _, num := range nums {
		freq[num]++
	}

	for num, count := range freq {
		if nextCount, ok := freq[num+1]; ok {
			maxLength = max(maxLength, count+nextCount)
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

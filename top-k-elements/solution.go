package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	unique := make([]int, 0, len(freqMap))
	for num := range freqMap {
		unique = append(unique, num)
	}

	n := len(unique)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if freqMap[unique[j]] > freqMap[unique[maxIdx]] {
				maxIdx = j
			}
		}
		unique[i], unique[maxIdx] = unique[maxIdx], unique[i]
	}

	return unique[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

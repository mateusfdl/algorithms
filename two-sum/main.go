package main

import "fmt"

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3, 4, 5, 6, 7, 8}, 11))
	fmt.Println(TwoSum2([]int{1, 2, 3, 4, 5, 6, 7, 8}, 11))

}

// O(n^2)
func TwoSum(list []int, key int) (int, int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if list[i]+list[j] == key {
				return i, j
			}

		}
	}

	return -1, -1
}

// O(n)
func TwoSum2(list []int, key int) (int, int) {
	m := make(map[int]int)

	for i := 0; i < len(list); i++ {
		target := key - list[i]

		if j, ok := m[target]; ok {
			return j, i
		}

		m[list[i]] = i
	}

	return -1, -1
}

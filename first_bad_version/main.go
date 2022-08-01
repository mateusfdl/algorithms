package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func firstBadVersion(n int) int {
	var left, right int
	left, right = 1, n

	for left < right {
		middle := left + (right-left)/2
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}

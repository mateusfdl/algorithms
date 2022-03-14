package main

import "fmt"

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}

func replaceElements(arr []int) []int {
	right := len(arr) - 1
	newNum := arr[len(arr)-1]
	maxNum := arr[len(arr)-1]

	for i := 0; i < len(arr)-1; i++ {
		if maxNum < arr[right-1] {
			maxNum = arr[right-1]
		}
		arr[right-1] = newNum

		newNum = maxNum

		right--
	}

	arr[len(arr)-1] = -1
	return arr
}

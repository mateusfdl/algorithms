package main

import "fmt"

type Case struct {
	n     int
	m     int
	nums1 []int
	nums2 []int
}

func main() {
	t := []Case{
		Case{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
		Case{
			nums1: []int{4, 5, 6, 0, 0, 0},
			nums2: []int{1, 2, 3},
			m:     3,
			n:     3,
		},
	}

	for _, c := range t {
		merge(c.nums1, c.m, c.nums2, c.n)
		fmt.Println(c.nums1)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		return
	}

	p1 := m - 1
	p2 := n - 1
	p := n + m - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

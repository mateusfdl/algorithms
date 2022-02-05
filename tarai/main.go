package main

import "fmt"

func main() {
	fmt.Println(tarai(10, 5, 0))
}

func tarai(x, y, z int) int {
	if y > x {
		return tarai(
			tarai(x-1, y, z),
			tarai(y-1, z, x),
			tarai(z-1, x, y),
		)
	} else {
		return y
	}
}

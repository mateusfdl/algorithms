package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	var value int

	for i := 0; i < len(s); i++ {
		currentValue := romanSymbols[s[i]]

		if i+1 < len(s) {
			nextValue := romanSymbols[s[i+1]]
			if currentValue < nextValue {
				value += nextValue - currentValue
				i++
				continue
			}
		}
		value += currentValue
	}
	return value
}

var romanSymbols = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

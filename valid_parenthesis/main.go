package main

import "fmt"

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []rune

	parser := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if parserBracketValue, ok := parser[v]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == parserBracketValue {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, v)
		}

	}

	return len(stack) == 0
}

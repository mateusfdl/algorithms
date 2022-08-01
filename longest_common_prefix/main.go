package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	var lowestWord string = strs[0]
	var prefix string

	for _, v := range strs {
		if len(lowestWord) > len(v) {
			lowestWord = v
		}
	}

	for i := 0; i < len(lowestWord); i++ {
		for _, s := range strs {
			if lowestWord[i] != s[i] {
				return prefix
			}
		}

		prefix += string(lowestWord[i])
	}

	return prefix
}

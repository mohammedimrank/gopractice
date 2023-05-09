package main

import "fmt"

func main() {

	s := "abcabcbb"

	m := make(map[byte]int)
	res := 0

	for l, r := 0, 0; r < len(s); r++ {

		if index, ok := m[s[r]]; ok {
			l = max(l, index)
		}

		res = max(res, r-l+1)

		m[s[r]] = r + 1

	}

	fmt.Println(res)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

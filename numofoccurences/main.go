package main

import "fmt"

func main() {

	nums := []int{1, 2}
	fmt.Println(uniqueElements(nums))
}

func uniqueElements(nums []int) bool {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	fmt.Println(m)

	freq := make(map[int]bool)

	for _, v := range m {
		fmt.Println(v)
		if freq[v] {
			return false
		}
		freq[v] = true
	}
	return true
}

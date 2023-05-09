package main

import "fmt"

func main() {

	nums := []int{8, 7, 2, 5, 3, 1}
	target := 9
	fmt.Println(getPairs(nums, target))

}

func getPairs(nums []int, target int) []int {

	pairs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := pairs[target-nums[i]]; ok {
			return []int{i, val}
		}

		pairs[nums[i]] = i
	}
	return []int{}

}

package main

import "fmt"

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum := nums[0]
	curSum := 0
	for _, val := range nums {

		if curSum < 0 {
			curSum = 0
		}
		curSum += val
		if maxSum < curSum {
			maxSum = curSum
		}
	}

	fmt.Println(maxSum)
}

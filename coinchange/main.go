package main

import (
	"fmt"
	"math"
)

func main() {

	num := 18
	coins := []int{1, 7, 5}

	dp := make([]int, num+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	ans := getCoins(coins, num, dp)
	fmt.Println(ans, dp)
}

func getCoins(arr []int, num int, dp []int) int {

	if num == 0 {
		return 0
	}
	ans := math.MaxInt

	for i := 0; i < len(arr); i++ {

		if num-arr[i] >= 0 {
			subAns := 0
			if dp[num-arr[i]] != -1 {
				subAns = dp[num-arr[i]]
			} else {
				subAns = getCoins(arr, num-arr[i], dp)
			}
			if subAns != math.MaxInt && subAns+1 < ans {
				ans = subAns + 1
			}
		}
	}
	dp[num] = ans
	return ans
}

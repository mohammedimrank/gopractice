package main

import "fmt"

func main() {

	fmt.Println(solution1(8))
	fmt.Println(solution2(8))
}

func solution1(n int) int {

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func solution2(n int) int {

	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		temp := one
		one = one + two
		two = temp
	}

	return one

}

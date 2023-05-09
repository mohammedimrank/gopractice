package main

import "fmt"

func main() {

	stocks := []int{7, 1, 5, 3, 6, 4}

	l := 0
	r := 1
	maxP := 0
	for r < len(stocks) {
		if stocks[l] < stocks[r] {
			val := stocks[r] - stocks[l]
			maxP = getMax(val, maxP)
		} else {
			l = r
		}
		r += 1
	}

	fmt.Println(maxP)
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

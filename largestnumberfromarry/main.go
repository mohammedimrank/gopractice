package main

import "fmt"

func main() {

	nums := []float64{5.2, 6, -9.8, 0.3, 5}

	for i := 1; i < len(nums); i++ {

		if nums[0] < nums[i] {
			nums[0] = nums[i]
		}
	}
	fmt.Println(nums[0])
	fmt.Println(swap(1, 2))
}

func swap(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

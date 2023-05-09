package main

import "fmt"

func main() {

	fmt.Println(findFiboNumber(0))
}

func findFiboNumber(n int) []int {

	fib1 := 0
	fib2 := 1
	res := []int{}
	inter := 0
	if n == 0 || n == 1 {
		return []int{1}
	}

	for i := 1; i <= n; i++ {
		inter = fib1 + fib2
		fib1 = fib2
		fib2 = inter
		res = append(res, inter)
	}
	return res
}

package main

import "fmt"

func checkPalindrome(n int) bool {

	temp := n
	sum := 0
	for n > 0 {

		remainder := n % 10
		sum = (sum * 10) + remainder
		n = n / 10
	}

	return temp == sum

}

func main() {

	fmt.Println(checkPalindrome(1222))

}

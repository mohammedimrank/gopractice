package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}

	set := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		set[arr[i]] = i
	}

	fmt.Println(len(set) != len(arr))
}

package main

import "fmt"

func main() {

	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println(arr)

	fmt.Println(reverseRecursive(arr))

	fmt.Println(reverseString("imraN"))
}

func reverseString(str string) string {

	byteArray := []rune(str)

	for i, j := 0, len(byteArray)-1; i < j; i, j = i+1, j-1 {
		byteArray[i], byteArray[j] = byteArray[j], byteArray[i]
	}
	return string(byteArray)
}

func reverseRecursive(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	return append(reverseRecursive(arr[1:]), arr[0])
}

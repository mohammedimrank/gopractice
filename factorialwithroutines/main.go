package main

import (
	"fmt"
	"sync"
)

func GetFactorial(n int) int {

	// fact := 1
	// for i := 1; i <= n; i++ {
	// 	fact = fact * i
	// }
	// return fact

	if n == 0 || n == 1 {
		return 1
	}
	return n * GetFactorial(n-1)
}

func main() {

	outputChan := make(chan int)
	inputList := []int{1, 2, 3, 4, 5, 6, 7}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for _, v := range inputList {
			val := GetFactorial(v)
			outputChan <- val
		}
		defer wg.Done()
		close(outputChan)
	}()

	go func() {
		defer wg.Done()
		for val := range outputChan {
			fmt.Println(val)
		}
	}()
	wg.Wait()

}

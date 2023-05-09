package main

import "fmt"

//not working hence deadlock
// func main() {
// 	c1 := make(chan int)
// 	fmt.Println("push c1: ")
// 	c1 <- 10
// 	g1 := <-c1
// 	fmt.Println("get g1: ", g1)
// }

func main() {

	c1 := make(chan int)
	go func() {
		g := <-c1
		fmt.Println("value of g is: ", g)
	}()

	c1 <- 10
	WithbufferedChan()
}
func WithbufferedChan() {

	c1 := make(chan int, 1)
	c1 <- 10
	g := <-c1
	fmt.Println("value of g is: ", g)
}

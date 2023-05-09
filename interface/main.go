package main

import "fmt"

type DanceOn interface {
	Dance()
}

type BreakDance struct {
	Name string
}

type AerialDance struct {
	Name string
}

func (b BreakDance) Dance() {

	fmt.Printf("%s is doing Breakdance on the floor\n", b.Name)
}

func (a AerialDance) Dance() {
	fmt.Printf("%s is doing Aerialdance on the air\n", a.Name)
}
func main() {
	b := BreakDance{
		Name: "break",
	}
	b.Dance()

	a := AerialDance{
		Name: "Aerial",
	}

	a.Dance()
}

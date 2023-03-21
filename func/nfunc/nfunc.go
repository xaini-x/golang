package main

import "fmt"

func greet(name string) {

	var displayName = func() {
		fmt.Println("Hi", name)
	}

	displayName()

}

func greet1() func() {
	return func() {
		fmt.Println("Hiiii xerx")
	}
}

func main() {

	greet("xaini")
	g1 := greet1()
	g1()
}

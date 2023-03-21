package main

import (
	"fmt"
)

func singleStruct() {

	type person struct {
		Fname string
		Lname string
		age   int
	}
	fmt.Println("struct work here")
	x := person{Fname: "xaini", Lname: "patel", age: 22}
	fmt.Println(x)
	fmt.Println(x.Fname)
}

type employee struct {
	Fname string
	Lname string
}
type office struct {
	employee
	id int
}

func inheritedStruct() {
	x := employee{"sushil", "patel"}
	y := office{employee{"xaini", "test"}, 1}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(y.employee)
}

func main() {
	singleStruct()
	inheritedStruct()
}

package main

import (
	"fmt"
)

func main() {

	fmt.Print("pointer started")

	var i = 400
	var p *int = &i

	fmt.Println("vlaue i 400:", i)

    fmt.Printf("Value of i in hexadecimal is %X\n", &i)

    fmt.Printf("Value of i in decimal is %v\n", &i)
    fmt.Printf("Value of p in decimal is %v\n", &p)
	fmt.Println("vlaue i location:", &i)
	fmt.Println("vlaue p location:", &p)
	fmt.Println("value p refer", *p)

}

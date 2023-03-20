package main

import "fmt"

func single() {
	fmt.Println("enter value ... ")
	var input int
	fmt.Scanln(&input)
	fmt.Println(" value ... ")
	for i := 0; i < input; i++ {
        fmt.Println(i)
    }

fmt.Println("forloop:")
}

func double() {
	fmt.Println("enter value ... ")
	var input int
	var input2 int
	fmt.Scanln(&input)
	fmt.Scanln(&input2)
	fmt.Println(" value ... ")
	for i := 0; i < input; i++ {
		for j := 0; j < input2; j++ {
            fmt.Println(j)
        }
        fmt.Println(i)
    }
}
//while loop type for loop
func forwhileloop(){
	fmt.Println("enter value... ")
    var input int
    // var input2 int
    fmt.Scanln(&input)
    // fmt.Scanln(&input2)
    fmt.Println(" value... ")
    for   input < 100  {
        input += input
        fmt.Println(input)
    }
}
func main() {
    single()
	double()
	forwhileloop()
}
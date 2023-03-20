package main

import "fmt"



func test() {
	var input int
	fmt.Print("enter your value... :")
	fmt.Scan(&input)
	if(input % 2 == 0){
		fmt.Printf("%v :is positve", input)
    }else{
        fmt.Printf("%v :is negative", input)
    }

}
func test1() {
	var i = 9
	fmt.Println(i)
	if(i % 2 == 0){
		fmt.Printf("%v :is positve", i)
    }else{
        fmt.Printf("%v :is negative\n", i)
    }

	}

	func main() {
		test()
		test1()
		fmt.Println("if-else , user input if-else")
	}

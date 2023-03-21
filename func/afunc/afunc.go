package main

import "fmt"
//normal afunc
var afunc = func() {
	fmt.Println("testing anonymus function")
}
//func with parameter and return
var anfunc = func(num1, num2 int) int {
	sum := num1 + num2

	return sum
}

func main() {
	afunc()
	sum := anfunc(10, 11)
	fmt.Println("sum of num :", sum)
}


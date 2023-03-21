package main

import "fmt"

func factorial(num int) int {

	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	num := 3
	var total int = factorial(num)
	fmt.Printf("factorial of total positive number %v is : %v",num, total)
}

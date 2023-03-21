package main

import "fmt"

func positive(num int) int {

	if num == 0 {
		return 0
	} else {
		return num + positive(num-1)
	}
}

func main() {
	var total int = positive(50)
	fmt.Println("sum of total positive number :", total)
}

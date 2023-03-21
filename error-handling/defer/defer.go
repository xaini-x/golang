package main

import ("fmt")

func main()  {

	fmt.Println("4")
	defer fmt.Println("1")
	fmt.Println("5")
	 defer fmt.Println("2")
	fmt.Println("6")
	defer fmt.Println("3")
}
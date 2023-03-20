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

// two for loops
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

// while loop type for loop
func forwhileloop() {
	fmt.Println("enter value... ")
	var input int
	// var input2 int
	fmt.Scanln(&input)
	// fmt.Scanln(&input2)
	fmt.Println(" value... ")
	for input < 100 {
		input += input
		fmt.Println(input)
	}
}

// break  type for loop
func breakloop() {
	var matches int
	fmt.Print("enter value... ")
	fmt.Scanln(&matches)

	for matches < 10 {
		fmt.Print("value is ", matches, "\n")
		matches++
		if matches > 5 {
			break
		}
	}

}

//continue loop type for loop
func continueInLoop(){
	var matches int
			fmt.Print("enter value... ")
			fmt.Scanln(&matches)
	
			for matches < 15{
				fmt.Print("value is ",matches,"\n")
					matches++
			if matches == 5{
				fmt.Println("thak gye hoge pani pee lo")
				matches +=5
				continue;

			}
				}
	
	}
func main() {
	single()
	double()
	forwhileloop()
	breakloop()
	continueInLoop()
}

package main

import ("fmt")

func oddNumber() func() int{
	num := 1

	return func() int {
		num = num + 2
		return num 
	}

}


func main (){
	odd := oddNumber()
	fmt.Println("odd number" , odd())
	fmt.Println("odd number" , odd())
	fmt.Println("odd number" , odd())
	fmt.Println("odd number" , odd())
	odd2 := oddNumber()
	fmt.Println("odd2 number" , odd2())
	fmt.Println("odd2 number" , odd2())
	fmt.Println("odd2 number" , odd2())
}

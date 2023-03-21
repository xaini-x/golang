package main

import ("fmt")

func divide(num1,num2 int)  {
	if num2 == 0 {
panic(" num 2 cant be 0")
	}else{
		result := num1/num2
		fmt.Println(result)
	}
}

func main()  {
	divide(4,2)
	divide(6,2)
	divide(4,0)
	divide(4,2)
}
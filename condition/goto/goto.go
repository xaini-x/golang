package main

import "fmt"

func main() {
	var age int
	Loop:
	fmt.Println("You are not eligible to vote ")  
	fmt.Print("enter the age")
	fmt.Scanf("%d", &age)
	if(age<= 17){
		goto Loop
   }else{
	fmt.Println(" eligible ")
}
	

}
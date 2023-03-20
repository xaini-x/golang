package main

import "fmt"

func switch_simple() {
fmt.Printf("enter number\n")
 var input int  
    fmt.Scanln(&input)  

    switch(input){
        case 1:
            fmt.Println("1")
        case 2:
            fmt.Println("2")
        case 3:
            fmt.Println("3")
        case 4:
            fmt.Println("4")
        case 5:
            fmt.Println("5")
        case 6:
            fmt.Println("6")
        case 7:
            fmt.Println("7")
        case 8:
            fmt.Println("8")
        case 9:
            fmt.Println("9")
        case 10:
            fmt.Println("10")
        case 11:
            fmt.Println("11")
        case 12:
            fmt.Println("12")
        case 13:
            fmt.Println("13")
        case 14:
            fmt.Println("14")
        default:
            fmt.Println("not implemented")
    }
}

func switch_advanced() {

fmt.Printf("enter number\n")
var input int  
   fmt.Scanln(&input)  

   switch(input){
	   case 1:
		//fall through will execute next condition without checking if input matches...
		   fmt.Println("1");fallthrough;
	   case 2:
		   fmt.Println("2");fallthrough;
	   case 3:
		   fmt.Println("3");fallthrough;
	   case 4:
		   fmt.Println("4");fallthrough;
	   case 5:
		   fmt.Println("5");fallthrough;
	   case 6:
		   fmt.Println("6");fallthrough;
	   case 7:
		   fmt.Println("7");fallthrough;
	   case 8:
		   fmt.Println("8");fallthrough;
	   case 9:
		   fmt.Println("9");fallthrough;
	   case 10:
		   fmt.Println("10");fallthrough;
	   case 11:
		   fmt.Println("11");fallthrough;
	   case 12:
		   fmt.Println("12");fallthrough;
	   case 13:
		   fmt.Println("13");fallthrough;
	   case 14:
		   fmt.Println("14");fallthrough;
	   default:
		   fmt.Println("not implemented")
   }
}
func main() {
switch_simple()
switch_advanced()
}
package main

import (
	"fmt"
)

func rec(sum int) {
	


	if sum > 0{
		fmt.Println(sum)
		rec(sum -1 )
		
	}else{
		fmt.Println("stop")
	}
	
}

func main(){
	rec(10)
}
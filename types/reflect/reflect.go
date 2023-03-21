package main

import ("fmt"
		"reflect"
		
)
func main() {
	var i string = "Testing Of Type Of "

	var j string = "testing of type Of "
	
	var k string = "Testing Of Type Of "
	//value of
	var l  = reflect.ValueOf([]string{"1","2","3"})
	var m = reflect.ValueOf([]string{"4","5","6"})
//deepequal check of string
fmt.Println("deep !equal :", reflect.DeepEqual(i,j))
fmt.Println("deep equal :", reflect.DeepEqual(i,k))
// copy value from one variable to another
fmt.Println("value before copy :", l)
fmt.Println("value before copy:", m)
fmt.Println("copy :", reflect.Copy(m,l))
fmt.Println("value after copy:", l)
fmt.Println("value after copy:", m)
// swapping value in array
var test = []int{1,2,3,4,5,6,7,8,9}
 swapValue :=reflect.Swapper(test)
fmt.Printf("orignal value %v",test)
swapValue(1,5)
fmt.Printf("after swap %v",test )
}
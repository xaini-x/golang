package main
import (
	"fmt"
	"strconv"
)

var i int = 1
var j string = "abcdefghijklmnop"
var k string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var l bool = false
var m float64 = 0.1
func simpleDataType() {
fmt.Printf("%T ,%T,%T,%T,%T\n", i, j, k, l, m)
fmt.Printf("%v ,%v,%q,%v,%v\n", i, j, k, l, m)
}

 
func typeCast() {  
  
   var i int = 10  
   var f float64 = 6.44  
   var str1 string = "101"  
   var str2 string = "10.123"  
  
   fmt.Println("int to float" ,float64(i))  
   fmt.Println("float to int" ,int(f))  
  
   newInt, _ := strconv.ParseInt(str1, 0, 64)  
   fmt.Println("str to int ",newInt)  
  
   newfloat, _ := strconv.ParseFloat(str2, 64)  
   fmt.Println("str to flaot",newfloat)  
}  

func main() {
	simpleDataType()
    typeCast()
}
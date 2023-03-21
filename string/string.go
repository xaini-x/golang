package main

import ("fmt"
		"reflect"
		"strings"
)
func main() {
// fmt.Print("string ")
var i string = "Testing Of Type Of "

 var n = []string{"a","b","c"}
var o string = "REPEAT "
var replace = "replace needed here"
fmt.Println("value :", i)
//type of variable
fmt.Println("type :" ,reflect.TypeOf(i))
fmt.Println("value :",reflect.ValueOf(i))

//toupper
fmt.Println("to uppercase ",strings.ToUpper(i))
//to lowecase
fmt.Println("to uppercase ",strings.ToLower(i))
//hasprefix
fmt.Println("has prefix ",strings.HasPrefix(i,"Te"))
// join array element
fmt.Println("join element",strings.Join(n,"*"))
// reaet value
fmt.Println("repeat element",strings.Repeat(o,4))
//replace value
fmt.Println("replace element" , strings.Replace(replace,"e","0",50))

}

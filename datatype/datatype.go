package main
import "fmt"

var i int = 1
var j string = "abcdefghijklmnop"
var k string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var l bool = false
var m float64 = 0.1
func main() {
fmt.Printf("%T ,%T,%T,%T,%T\n", i, j, k, l, m)
fmt.Printf("%v ,%v,%q,%v,%v\n", i, j, k, l, m)
}
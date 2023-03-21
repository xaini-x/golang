package main

import "fmt"

func singleArray() {
 var arr = []int{1, 2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17}
 fmt.Println("single arrays")
 for i := 0; i <len(arr);i++ {
	arr[i] = arr[i]+10
	
}
fmt.Println(arr)
}

func multiDimensionalArray() {
	var arr =[3][3]int{ {1,2,3}, {4,5,6} ,{7,8,9}}
fmt.Println("multidimenssional arrays")
    for i := 0; i < len(arr); i++ {
       for j := 0; j < len(arr[i]); j++ {
		fmt.Print(arr[i][j])
	}
    fmt.Println()

}
}

func main() {
    singleArray()
    multiDimensionalArray()
}
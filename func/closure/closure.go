package main

import (
	"fmt"
)

func greet() func() string {
	name := "xaini"

	return func() string {
		name = "hii " + name
		return name

	}

}

func main() {
	message := greet()
	fmt.Println(message())
}

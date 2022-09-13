package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var x int = 20
	var y float64 = 2.52
	var z float64 = float64(x) * y
	var d int = x * int(z)

	fmt.Println(z, d)
}

package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		// shadow using := , just = will print 5 at the final println()
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

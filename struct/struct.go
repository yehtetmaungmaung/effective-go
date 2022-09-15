package main

import "fmt"

func main() {
	fmt.Println(("=========="))
	type Person struct {
		name string
		age  int
	}
	f := Person{
		name: "Bob",
		age:  22,
	}
	var g struct {
		name string
		age  int
	}
	g = f
	fmt.Println(f == g)
}

package main

import "fmt"

func main() {
	fmt.Println("----------------")

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	fmt.Println("----------------")
	v, ok = m["world"]
	fmt.Println(v, ok)

	fmt.Println("----------------")
	v, ok = m["dogs"]
	fmt.Println(v, ok)

	fmt.Println("----------------")
	delete(m, "hello")
	v, ok = m["hello"]
	fmt.Println(v, ok)
}

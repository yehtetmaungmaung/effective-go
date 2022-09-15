package main

import "fmt"

func main() {
	fmt.Println("-----------")
	m := map[int]bool{}
	vals := []int{2, 4, 5, 3, 6, 8, 7, 9, 10, 1}
	for _, v := range vals {
		m[v] = true
	}
	fmt.Println(len(vals), len(m))
	fmt.Println(m[5])
	fmt.Println(m[344])
	if m[100] {
		fmt.Println("100 is in set")
	} else {
		fmt.Println("100 is not in set")
	}
}

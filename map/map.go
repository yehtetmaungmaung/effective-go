package main

import (
	"fmt"
)

func main() {
	fmt.Println("program started ..")
	totalWins := map[string]int{}
	totalWins["orcas"] = 1
	totalWins["lions"] = 2
	fmt.Println(totalWins["orcas"])
	fmt.Println(totalWins["lions"])
	fmt.Println("----------")
	totalWins["orcas"]++
	totalWins["lions"] = 3
	fmt.Println(totalWins["orcas"])
	fmt.Println(totalWins["lions"])
	fmt.Println("cats", totalWins["cats"])
	fmt.Println("----------")
	totalWins["cats"]++
	fmt.Println("cats", totalWins["cats"])
}

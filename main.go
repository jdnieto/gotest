package main

import (
	"fmt"
)

func main() {
	scores := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		scores = append(scores, 5)
	}
	test := make([]int, 5)
	test = append(test, 5)
	fmt.Println(scores)
	fmt.Println(test)
}

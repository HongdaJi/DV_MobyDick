package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)
	b := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b[i])
	}
	c := make([]int, a)
	for i := 0; i < a; i++ {
		for j := i; j > -1; j-- {
			if b[j] < b[i] {
				c[i]++
			}
		}
	}
	for i := 0; i < a; i++ {
		fmt.Print(c[i], " ")
	}

}

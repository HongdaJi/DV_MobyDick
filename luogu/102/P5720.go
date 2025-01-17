package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	i := 2
	for ; a != 1; i++ {
		a /= 2
	}
	fmt.Print(i - 1)
}

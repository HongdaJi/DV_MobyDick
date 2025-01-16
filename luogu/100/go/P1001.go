package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Scanln(&num1, &num2)
	b := num1 + num2
	fmt.Println(b)
}

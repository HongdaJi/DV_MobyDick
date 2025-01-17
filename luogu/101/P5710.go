package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)

	one := a%2 == 0

	two := a > 4 && a < 13

	if one && two {
		fmt.Print(1)

	} else {
		fmt.Print(0)
	}
	fmt.Print(" ")

	if one || two {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
	fmt.Print(" ")

	if one == true && two == false || one == false && two == true {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
	fmt.Print(" ")

	if one == false && two == false {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}

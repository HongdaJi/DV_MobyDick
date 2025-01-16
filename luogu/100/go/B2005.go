package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Print(`  `, input, "\n")
	fmt.Print(` `, input, input, input, "\n")
	fmt.Print(input, input, input, input, input, "\n")
}

package main

import (
	"fmt"
)

func main() {
	var a float32
	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		return
	}

	b := float32(2)
	c := float32(2)
	i := 1

	for c <= a {
		c += b * 0.98
		b *= 0.98
		i++
	}

	fmt.Print(i)
}

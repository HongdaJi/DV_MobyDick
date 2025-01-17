package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scanln(&a)
	b := make([][]int, a)
	sum := make([]int, a)
	var num int
	for i := range b {
		b[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Scan(&b[i][j])
		}
		sum[i] = b[i][0] + b[i][1] + b[i][2]
	}

	for i := 0; i < a; i++ {

		for j := i + 1; j < a; j++ {
			if math.Abs(float64(b[i][1]-b[j][1])) <= 5 &&
				math.Abs(float64(b[i][2]-b[j][2])) <= 5 &&
				math.Abs(float64(b[i][0]-b[j][0])) <= 5 &&
				math.Abs(float64(sum[i]-sum[j])) <= 10 {
				num++
			}
		}
	}
	fmt.Print(num)
}

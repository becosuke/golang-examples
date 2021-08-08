package main

import (
	"fmt"
)

func main() {
	// 2640
	// max := 10
	max := 100
	min := 1
	powsum := 0
	sum := 0
	for i := min; i <= max; i++ {
		powsum += i * i
		sum += i
	}
	sumpow := sum * sum
	fmt.Println(sumpow - powsum)
}

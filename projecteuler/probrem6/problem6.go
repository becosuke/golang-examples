package main

import "fmt"

func main() {
	// 2640
	// max := 10
	max := 100
	min := 1
	powSum := 0
	sum := 0
	for i := min; i <= max; i++ {
		powSum += i * i
		sum += i
	}
	sumPow := sum * sum
	fmt.Println(sumPow - powSum)
}

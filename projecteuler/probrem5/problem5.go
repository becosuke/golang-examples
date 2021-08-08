package main

import "fmt"

func main() {
	// 2520
	// max := 10
	max := 20
	min := 1
	result := min
	for i := min + 1; i <= max; i++ {
		result = lcm(result, i)
	}
	fmt.Println(result)
}

// greatest common divisor
func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

// least common multiple
func lcm(x int, y int) int {
	return x * y / gcd(x, y)
}

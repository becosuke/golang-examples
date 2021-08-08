package main

import "fmt"

func main() {
	// n := 13195
	n := 600851475143
	result := 0
	i := 2
	for n > 1 {
		if n%i == 0 {
			result = i
			n = n / i
			i = 2
		} else {
			i += 1
		}
	}
	fmt.Println(result)
}

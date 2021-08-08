package main

import (
	"fmt"
)

func main() {
	before := 1
	current := 2
	result := 0
	for current < 4000000 {
		before, current = current, before+current
		if before%2 == 0 {
			result += before
		}
	}
	fmt.Println(result)
}

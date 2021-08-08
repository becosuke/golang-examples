package main

import (
	"fmt"
)

func main() {
	// 9009
	// max := 99
	// min := 10
	max := 999
	min := 100
	result := 0
	for i := max; i >= min; i-- {
		if i*max < result {
			break
		}
		for j := max; j >= min; j-- {
			n := i * j
			if n < result {
				break
			}
			r := 0
			tmp := n
			for tmp >= 1 {
				r = tmp%10 + r*10
				tmp = tmp / 10
			}
			if n == r {
				result = n
			}
		}
	}
	fmt.Println(result)
}

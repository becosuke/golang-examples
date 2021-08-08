package main

import "fmt"

func main() {
	// 6 -> 2, 3, 5, 7, 11, 13
	// cnt := 6
	cnt := 10001
	min := 2
	primes := make([]int, 0, cnt)
	for i := min; len(primes) < cnt; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes[len(primes)-1])
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

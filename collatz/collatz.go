package collatz

func calc(num int) int {
	if num%2 == 0 {
		return num / 2
	} else {
		return num*3 + 1
	}
}

func Collatz(num int) []int {
	results := []int{}
	for {
		results = append(results, num)
		if num == 1 {
			break
		}
		num = calc(num)
	}
	return results
}

package etc

import "fmt"

func SW(n int) {
	switch n {
	case 0:
	case 1:
		fmt.Printf("n is %d\n", n)
	case 2:
		fmt.Print("n is two\n")
	default:
		fmt.Print("unknown\n")
	}
}

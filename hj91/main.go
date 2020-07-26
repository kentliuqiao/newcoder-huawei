package main

import "fmt"

func main() {
	var m, n int
	for {
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			break
		}
		fmt.Println(steps(m, n))
	}
}

func steps(m, n int) int {
	if m == 0 || n == 0 {
		return 1
	}

	return steps(m-1, n) + steps(m, n-1)
}

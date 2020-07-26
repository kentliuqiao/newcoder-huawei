package main

import "fmt"

func main() {
	var s1, s2 string
	for {
		n, err := fmt.Scan(&s1, &s2)
		if n < 2 || err != nil {
			return
		}

	}
}

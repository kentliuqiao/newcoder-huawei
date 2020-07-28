package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2 string
	for {
		n, err := fmt.Scan(&s1, &s2)
		if n < 2 || err != nil {
			return
		}

		max := 0
		for i := 0; i < len(s1); i++ {
			for j := i + 1; j <= len(s1); j++ {
				if strings.Contains(s2, s1[i:j]) && j-i > max {
					max = j - i
				}
			}
		}
		fmt.Println(max)
	}
}

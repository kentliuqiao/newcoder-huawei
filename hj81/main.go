package main

import (
	"fmt"
	"strings"
)

func main() {
	var short, long string
	for {
		if _, err := fmt.Scan(&short); err != nil {
			return
		}
		fmt.Scan(&long)

		flag := true
		for _, b := range short {
			if !strings.Contains(long, string(b)) {
				flag = false
			}
		}

		fmt.Println(flag)
	}
}
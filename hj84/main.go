package main

import "fmt"

func main() {
	var s string
	for {
		if _, err := fmt.Scan(&s); err != nil {
			return
		}

		cnt := 0
		for _, b := range s {
			if b >= 'A' && b <= 'Z' {
				cnt++
			}
		}

		fmt.Println(cnt)
	}
}

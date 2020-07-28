package main

import (
	"fmt"
)

func main() {
	for {
		str := ""
		if _, err := fmt.Scan(&str); err != nil {
			break
		}

		res := make([]int, 256)
		for _, s := range str {
			res[s]++
		}

		max := 0
		for {
			for i := range res {
				if res[i] == 0 {
					continue
				}
				if res[i] > res[max] {
					max = i
				}
			}
			if res[max] == 0 {
				break
			}

			res[max] = 0
			fmt.Printf("%c", max)
		}
		fmt.Println()
	}
}

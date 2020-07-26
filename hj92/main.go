package main

import "fmt"

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}

		cnt, max := 0, 0
		dp := make([]int, len(str))
		for i, s := range str {
			if s >= '0' && s <= '9' {
				cnt++
				if max < cnt {
					max = cnt
				}
			} else {
				cnt = 0
			}
			dp[i] = cnt
		}

		for i, v := range dp {
			if v == max {
				fmt.Printf("%s", str[i-max+1:i+1])
			}
		}
		fmt.Printf(",%d\n", max)
	}
}

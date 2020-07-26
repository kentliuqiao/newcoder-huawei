package main

import "fmt"

func main() {
	var s string
	for {
		if _, err := fmt.Scan(&s); err != nil {
			return
		}

		var s1 string
		for i := len(s) - 1; i >= 0; i-- {
			s1 += string(s[i])
		}

		max := 0
		dp := make([]int, len(s))

		for j := 0; j < len(s); j++ {
			for i := len(s) - 1; i >= 0; i-- {
				if s[j] == s1[i] {
					if i >= 1 {
						dp[i] = dp[i-1] + 1
					} else {
						dp[i] = 1
					}
					if max < dp[i] {
						max = dp[i]
					}
				} else {
					dp[i] = 0
				}
			}
		}

		fmt.Println(max)

	}
}

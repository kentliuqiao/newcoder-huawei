package main

import "fmt"

func main() {
	var (
		wild string
		str  string
	)

	for {
		if _, err := fmt.Scan(&wild, &str); err != nil {
			return
		}

		l1 := len(wild)
		l2 := len(str)

		dp := make([][]bool, l1+1)
		for i := range dp {
			dp[i] = make([]bool, l2+1)
		}

		dp[0][0] = true

		for i := 1; i <= l2; i++ {
			dp[0][i] = false
		}
		for i := 1; i <= l1; i++ {
			dp[i][0] = dp[i-1][0] && wild[i] == '*'
		}

		for i := 1; i <= l1; i++ {
			for j := 1; j <= l2; j++ {
				w := wild[i-1]
				s := str[j-1]
				if w == '*' {
					dp[i][j] = dp[i-1][j] || dp[i][j-1]
				} else {
					dp[i][j] = (w == s || w == '?') && dp[i-1][j-1]
				}
			}
		}

		fmt.Println(dp[l1][l2])
	}
}

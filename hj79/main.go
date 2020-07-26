package main

import (
	"fmt"
)

const ic = 1
const dc = 1
const rc = 1

func min(a, b, c int) int {
	re := a
	if b < re {
		re = b
	}
	if c < re {
		re = c
	}
	return re
}

func main() {
	var (
		str1 string
		str2 string
	)

	for {
		_, err := fmt.Scan(&str1)
		if err != nil {
			break
		}
		fmt.Scan(&str2)

		if str1 == str2 {
			fmt.Println(0)
			continue
		}
		if str1 == "" {
			fmt.Println(len(str2) * rc)
			continue
		}
		if str2 == "" {
			fmt.Println(len(str1) * dc)
			continue
		}

		l1, l2 := len(str1), len(str2)
		dp := make([][]int, l1+1)
		for ix := range dp {
			dp[ix] = make([]int, l2+1)
		}

		for i := 0; i <= l1; i++ {
			// 将字符串 变成 空串
			dp[i][0] = i * dc
		}
		for j := 0; j <= l2; j++ {
			// 将空串 变成 字符串
			dp[0][j] = j * ic
		}

		for i := 1; i <= l1; i++ {
			for j := 1; j <= l2; j++ {
				case1 := dp[i-1][j] + dc // 删除一个字符后变成 str2
				case2 := dp[i][j-1] + ic // 插入一个字符后变成 str2
				case3 := 0
				if str1[i-1] == str2[j-1] { // 从空串算起，字符串的索引要减 1
					case3 = dp[i-1][j-1]
				} else {
					case3 = dp[i-1][j-1] + rc // 修改一个字符变成 str2
				}
				dp[i][j] = min(case1, case2, case3)
			}
		}
		fmt.Printf("1/%d\n", dp[l1][l2]+1)
	}
}

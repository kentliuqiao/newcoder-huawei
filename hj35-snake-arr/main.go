package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		if _, err := fmt.Scan(&n); err != nil {
			return
		}

		arr := make([][]int, n)
		for i := range arr {
			arr[i] = make([]int, n)
		}

		start := 1
		for k := 0; k < n; k++ {
			for i, j := k, 0; i >= 0 && j <= k; {
				arr[i][j] = start
				start++
				i--
				j++
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i][j] != 0 {
					fmt.Printf("%d ", arr[i][j])
				}
			}
			fmt.Println()
		}

	}
}

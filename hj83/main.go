package main

import "fmt"

func main() {
	var m, n int
	for {
		if _, err := fmt.Scan(&m, &n); err != nil {
			return
		}
		if m > 9 || n > 9 {
			fmt.Println(-1)
		}

		table := make([][]int, m)
		for i := range table {
			table[i] = make([]int, n)
		}
		fmt.Println(0)

		x1, y1 := 0, 0
		x2, y2 := 0, 0
		fmt.Scan(&x1, &y1, &x2, &y2)
		if x1 > m-1 || y1 > n-1 || x2 > m-1 || y2 > n-1 {
			fmt.Println(-1)
		} else {
			fmt.Println(0)
		}

		row := 0
		fmt.Scan(&row)
		if row < 0 || row > m || m == 9 {
			fmt.Println(-1)
		} else {
			fmt.Println(0)
		}

		col := 0
		fmt.Scan(&col)
		if col < 0 || col > n || n == 9 {
			fmt.Println(-1)
		} else {
			fmt.Println(0)
		}

		x, y := 0, 0
		fmt.Scan(&x, &y)
		if x >= m || y >= n {
			fmt.Println(-1)
		} else {
			fmt.Println(0)
		}

	}
}

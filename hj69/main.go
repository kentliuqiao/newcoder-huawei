package main

import (
	"fmt"
)

func main() {
	for {
		r, rc, c := 0, 0, 0
		if _, err := fmt.Scan(&r, &rc, &c); err != nil {
			return
		}

		arr1 := make([][]int, r)
		for i := 0; i < r; i++ {
			arr1[i] = make([]int, rc)
		}
		arr2 := make([][]int, rc)
		for i := 0; i < rc; i++ {
			arr2[i] = make([]int, c)
		}
		for i := 0; i < r; i++ {
			for j := 0; j < rc; j++ {
				fmt.Scan(&arr1[i][j])
			}
		}
		for i := 0; i < rc; i++ {
			for j := 0; j < c; j++ {
				fmt.Scan(&arr2[i][j])
			}
		}

		res := make([][]int, r)
		for i := 0; i < r; i++ {
			res[i] = make([]int, c)
		}

		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				for k := 0; k < rc; k++ {
					res[i][j] += arr1[i][k] * arr2[k][j]
				}
			}
		}

		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				fmt.Printf("%d ", res[i][j])
			}
			fmt.Println()
		}
	}
}

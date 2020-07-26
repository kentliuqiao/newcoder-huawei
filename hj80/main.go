package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2 int
	for {
		if _, err := fmt.Scan(&n1); err != nil {
			return
		}

		arr1 := make([]int, n1)
		for i := 0; i < n1; i++ {
			num := 0
			fmt.Scan(&num)
			arr1[i] = num
		}

		fmt.Scan(&n2)

		arr2 := make([]int, n2)
		for i := 0; i < n2; i++ {
			num := 0
			fmt.Scan(&num)
			arr2[i] = num
		}
		sort.Ints(arr2)
		sort.Ints(arr1)

		res := make([]int, 0)
		i, j := 0, 0
		for i < len(arr1) && j < len(arr2) {
			if arr1[i] < arr2[j] {
				res = append(res, arr1[i])
				i++
			} else if arr1[i] > arr2[j] {
				res = append(res, arr2[j])
				j++
			} else {
				i++
			}
		}

		res = append(res, arr1[i:]...)
		res = append(res, arr2[j:]...)

		for _, v := range res {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

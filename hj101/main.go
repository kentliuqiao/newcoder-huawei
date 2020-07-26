package main

import (
	"fmt"
	"sort"
)

func main() {
	order()
}

func order() {
	var (
		n    int
		flag int
		tmp  int
	)

	for {
		num, err := fmt.Scan(&n)
		if num <= 0 || err != nil {
			break
		}
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&tmp)
			nums[i] = tmp
		}

		fmt.Scan(&flag)

		sort.Ints(nums)

		if flag == 1 {
			for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

		for i := range nums {
			if i == 0 {
				fmt.Print(nums[i])
			} else {
				fmt.Printf(" %d", nums[i])
			}

		}

		fmt.Println()
	}
}

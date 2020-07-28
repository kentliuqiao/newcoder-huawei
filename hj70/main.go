package main

import "fmt"

func main() {
	for {
		n := 0
		if _, err := fmt.Scan(&n); err != nil {
			return
		}
		nums := make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&nums[i][0], &nums[i][1])
		}

		s := ""
		fmt.Scan(&s)
		var res [][]int
		ans, idx := 0, 0

		for _, v := range s {
			if v == ')' && len(res) >= 2 {
				num1 := res[len(res)-1]
				num2 := res[len(res)-2]
				res = res[0 : len(res)-2]
				ans += num2[0] * num2[1] * num1[1]
				num3 := []int{num2[0], num1[1]}
				res = append(res, num3)
			} else if v >= 'A' && v <= 'Z' {
				num3 := []int{nums[idx][0], nums[idx][1]}
				res = append(res, num3)
				idx++
			}
		}

		fmt.Println(ans)
	}
}

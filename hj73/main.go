package main

import "fmt"

func main() {
	for {
		y, m, d := 0, 0, 0
		_, err := fmt.Scan(&y)
		if err != nil {
			return
		}
		_, err = fmt.Scan(&m)
		if err != nil {
			return
		}
		_, err = fmt.Scan(&d)
		if err != nil {
			return
		}

		days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		if year(y) {
			days[1] = 29
		}
		res := 0
		for i := 0; i < m-1; i++ {
			res += days[i]
		}
		fmt.Println(res + d)
	}
}

func year(y int) bool {
	if y%4 == 0 && y%100 != 0 || y%400 == 0 {
		return true
	}
	return false
}

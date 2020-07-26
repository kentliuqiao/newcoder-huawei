package main

import "fmt"

func main() {
	var num int
	for {
		if _, err := fmt.Scan(&num); err != nil {
			return
		}
		tm := num*num - num + 1
		for i := 0; i < num; i++ {
			if i == 0 {
				fmt.Print(tm)
			} else {
				fmt.Printf("+%d", tm)
			}
			tm += 2
		}
		fmt.Println()
	}
}

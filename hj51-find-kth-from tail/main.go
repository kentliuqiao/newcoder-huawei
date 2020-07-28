package main

import "fmt"

type node struct {
	Val  int
	Next *node
}

func main() {
	for {
		var total, k int
		if _, err := fmt.Scan(&total); err != nil {
			return
		}
		if total == 0 {
			return
		}
		dummy := &node{}
		p := dummy
		tmp := 0
		for i := 0; i < total; i++ {
			fmt.Scan(&tmp)
			p.Next = &node{Val: tmp}
			p = p.Next
		}
		fmt.Scan(&k)
		if k > total || k < 0 {
			fmt.Println(0)
			continue
		}

		p = dummy
		for i := 0; i <= total-k; i++ {
			p = p.Next
		}

		fmt.Println(p.Val)
	}
}

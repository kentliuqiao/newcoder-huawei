package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		var it, rt int
		if _, err := fmt.Scan(&it); err != nil {
			return
		}
		I := make([]int, it)
		for i := 0; i < it; i++ {
			fmt.Scan(&I[i])
		}

		fmt.Scan(&rt)
		R := make([]int, rt)
		rMap := make(map[int]bool)
		for i := 0; i < rt; i++ {
			tmp := 0
			fmt.Scan(&tmp)
			R[i] = tmp
			rMap[tmp] = true
		}
		R1 := make([]int, 0)
		for _, v := range R {
			if rMap[v] == true {
				R1 = append(R1, v)
				rMap[v] = false
			}
		}

		sort.Ints(R1)

		res := make([]int, 0)
		for _, r := range R1 {
			tmp := make([]int, 0)
			flag := false
			rstr := strconv.Itoa(r)
			for i, v := range I {
				vstr := strconv.Itoa(v)
				if strings.Contains(vstr, rstr) {
					flag = true
					tmp = append(tmp, i)
					tmp = append(tmp, v)
				}
			}
			if flag {
				res = append(res, r)
				res = append(res, len(tmp)/2)
				res = append(res, tmp...)
			}
		}

		fmt.Printf("%d ", len(res))
		for _, v := range res {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

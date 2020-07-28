package main

import "fmt"

func main() {
	var s1, s2 string
	for {
		if _, err := fmt.Scan(&s1, &s2); err != nil {
			return
		}

		l1, l2 := len(s1)+1, len(s2)+1
		arr := make([][]int, l1)
		for i := range arr {
			arr[i] = make([]int, l2)
		}
		for i := 0; i < l1; i++ {
			arr[i][0] = i
		}
		for i := 0; i < l2; i++ {
			arr[0][i] = i
		}
		for i := 1; i < l1; i++ {
			for j := 1; j < l2; j++ {
				tmp := 0
				if s1[i-1] == s2[j-1] {
					tmp = arr[i-1][j-1]
				} else {
					tmp = arr[i-1][j-1] + 1
				}
				arr[i][j] = min(tmp, arr[i][j-1]+1, arr[i-1][j]+1)
			}
		}
		fmt.Printf("1/%d\n", arr[l1-1][l2-1]+1)
	}
}

func min(x, y, z int) int {
	t1 := 0
	if x < y {
		t1 = x
	} else {
		t1 = y
	}
	t2 := 0
	if t1 < z {
		t2 = t1
	} else {
		t2 = z
	}
	return t2
}

package main

import (
	"fmt"
	"strconv"
)

func mydd(n1, n2 string) string {
	tmp := ""
	// 让n1永远是小 n2大
	if len(n1) > len(n2) {
		tmp = n1
		n1 = n2
		n2 = tmp
	}
	n1l, n2l := len(n1), len(n2)
	cn := n2l - n1l
	pad := 0 // 借值
	r := ""

	for i := n1l - 1; i >= 0; i-- {
		n11, _ := strconv.Atoi(string(n1[i]))
		n21, _ := strconv.Atoi(string(n2[i+cn]))

		newV := n11 + n21 + pad
		pad = 0
		if newV >= 10 {
			pad = 1
			newV -= 10
		}
		r = fmt.Sprintf("%d", newV) + r
	}

	for i := cn - 1; i >= 0; i-- {
		if pad != 0 {
			n21, _ := strconv.Atoi(string(n2[i]))
			newV := n21 + pad
			pad = 0
			if newV >= 10 {
				pad = 1
				newV -= 10
			}
			r = fmt.Sprintf("%d", newV) + r

			continue
		}
		r = string(n2[i]) + r
	}
	if pad != 0 {
		r = fmt.Sprintf("%d", pad) + r

	}
	return r
}
func main() {
	for {

		n1, n2 := "", ""
		if n, _ := fmt.Scanf("%s\n", &n1); n == 0 || len(n1) == 0 {
			break
		}
		if n, _ := fmt.Scanf("%s\n", &n2); n == 0 || len(n2) == 0 {
			break
		}
		fmt.Printf("%v\n", mydd(n1, n2))

	}

}

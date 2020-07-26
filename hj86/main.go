package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	for {
		if _, err := fmt.Scan(&s); err != nil {
			return
		}
		i, _ := strconv.Atoi(s)
		fmt.Println(maxContinueBit(i))
	}
}

func maxContinueBit(s int) int {
	k := 0
	for ; s != 0; k++ {
		s = s & (s << 1)
	}
	return k
}

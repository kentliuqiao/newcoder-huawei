package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	for {
		if _, err := fmt.Scan(&n); err != nil {
			break
		}
		p, q, r := big.NewInt(0), big.NewInt(0), big.NewInt(1)
		for i := int64(0); i < n; i++ {
			p.Set(q)
			q.Set(r)
			r = r.Add(q, p)
		}
		fmt.Println(r)
	}
}

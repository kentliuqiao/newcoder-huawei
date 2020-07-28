package main

import (
	"fmt"
)

func main() {
	num := 0.0
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		last := 1.0
		new := last - (last*last*last-num)/(3*last*last)
		for (last*last*last-num) > 0.00001 ||
			(last*last*last-num) < -0.00001 {
			last = new
			new = last - (last*last*last-num)/(3*last*last)
		}
		fmt.Printf("%.1f\n", last)
	}
}

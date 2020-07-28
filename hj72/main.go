package main

import "fmt"

func main() {
	var m int
	for {
		if _, err := fmt.Scan(&m); err != nil {
			return
		}

		for x := 0; x <= 20; x += 4 {
			y := (100 - 7*x) / 4
			z := 75 + 3*x/4
			if y > 0 && z > 0 && x+y+z == 100 {
				fmt.Println(x, y, z)
			}
		}
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	for {
		if _, err := fmt.Scan(&str); err != nil {
			return
		}

		res := ""
		for i := range str {
			if str[i] >= '0' && str[i] <= '9' {
				res += "*" + string(str[i]) + "*"
			} else {
				res += string(str[i])
			}
		}

		res = strings.ReplaceAll(res, "**", "")
		fmt.Println(res)
	}
}

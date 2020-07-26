package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ip string
	for {
		_, err := fmt.Scan(&ip)
		if err != nil {
			break
		}

		fmt.Println(isValid(ip))
	}
}

func isValid(ip string) string {
	if len(ip) == 0 {
		return "YES"
	}

	strs := strings.Split(ip, ".")
	if len(strs) < 4 {
		return "NO"
	}

	for _, str := range strs {
		val, err := strconv.Atoi(str)
		if err != nil {
			return "NO"
		}

		if val < 0 || val > 255 || (val != 0 && strs[0] == "0") {
			return "NO"
		}
	}

	return "YES"
}

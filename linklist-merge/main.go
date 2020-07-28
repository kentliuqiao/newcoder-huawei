package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str1 := scanner.Text()
	list1 := make([]int, 0)
	for _, v := range strings.Split(str1, " ") {
		tmp, _ := strconv.Atoi(v)
		list1 = append(list1, tmp)
	}

	scanner.Scan()
	str2 := scanner.Text()
	list2 := make([]int, 0)
	for _, v := range strings.Split(str2, " ") {
		tmp, _ := strconv.Atoi(v)
		list2 = append(list2, tmp)
	}

	p, q := 0, 0
	for i := 0; i < len(list1)+len(list2); i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		if p == len(list1) {
			fmt.Print(list2[q])
			q++
		} else if q == len(list2) {
			fmt.Print(list1[p])
			p++
		} else {
			if list1[p] < list2[q] {
				fmt.Print(list1[p])
				p++
			} else {
				fmt.Print(list2[q])
				q++
			}
		}
	}
	fmt.Println()
}

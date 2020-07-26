package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pattern := "345678910JQKA2jokerJOKER"

	for {
		reader := bufio.NewReader(os.Stdin)
		s, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		s1 := string(s)
		if len(s1) < 3 {
			return
		}

		if strings.Contains(s1, "joker") {
			fmt.Println("joker JOKER")
			continue
		}
		cards := strings.Split(s1, "-")

		ts0 := strings.Split(cards[0], " ")
		ts1 := strings.Split(cards[1], " ")

		if len(ts0) != len(ts1) {
			if len(ts0) == 4 {
				fmt.Println(cards[0])
				continue
			} else if len(ts1) == 4 {
				fmt.Println(cards[1])
				continue
			} else {
				fmt.Println("ERROR")
				continue
			}
		}

		i0 := strings.LastIndex(pattern, ts0[0])
		i1 := strings.LastIndex(pattern, ts1[0])
		if i0 < i1 {
			fmt.Println(cards[1])
		} else {
			fmt.Println(cards[0])

		}
	}
}

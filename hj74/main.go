package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	out := strings.Fields(str)

	fmt.Println(len(out))

	for _, v := range out {
		if v[0] == '"' {
			fmt.Println(v[1 : len(v)-1])
		} else {
			fmt.Println(v)
		}
	}
}

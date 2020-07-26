package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func train(inputs, stack []int, outputs []string, res *[]string, index int) {
	// fmt.Println("---", inputs[index:], "---")
	if index == len(inputs) {
		for len(stack) > 0 {
			outputs = append(outputs, strconv.Itoa(stack[len(stack)-1]))
			stack = stack[:len(stack)-1]
		}

		// fmt.Println("outputs:", outputs)
		str := strings.Join(outputs, " ")
		*res = append(*res, str)
		return
	}

	// fmt.Println("push:", inputs[index])
	stackCopy := make([]int, len(stack)+1)
	copy(stackCopy, stack)
	stackCopy[len(stack)] = inputs[index]
	train(inputs, stackCopy, outputs, res, index+1)

	// 1 2 3 5 6 7 4
	if len(stack) > 0 {
		// fmt.Println("[pop]", stack, "index:", index, "outputs:", outputs)
		outputs = append(outputs, strconv.Itoa(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
		train(inputs, stack, outputs, res, index)
	}
}

func main() {
	for {
		var count int
		_, err := fmt.Scan(&count)
		if err != nil {
			break
		}
		inputs := make([]int, count)
		for i := 0; i < count; i++ {
			fmt.Scan(&(inputs[i]))
		}
		res := make([]string, 0)
		train(inputs, make([]int, 0), make([]string, 0), &res, 0)
		sort.Strings(res)
		for _, n := range res {
			fmt.Println(n)
		}
	}
}

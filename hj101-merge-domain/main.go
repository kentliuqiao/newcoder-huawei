package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type domain struct {
	start, end int
}

func (d domain) String() string {
	return fmt.Sprintf("%d,%d", d.start, d.end)
}

func (d domain) Less(o domain) bool {
	return d.end < o.end
}

type sortBy []domain

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i].Less(a[j]) }

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _, err := reader.ReadLine()
		if err != nil {
			return
		}

		temp := strings.Split(string(str), " ")

		arr := make(sortBy, 0)

		for _, v := range temp {
			strs := strings.Split(v, ",")
			start, _ := strconv.Atoi(strs[0])
			end, _ := strconv.Atoi(strs[1])
			arr = append(arr, domain{start: start, end: end})
		}

		sort.Sort(arr)

		if len(arr) < 1 {
			break
		}

		tmp := domain{arr[0].start, arr[0].end}
		for i := 1; i < arr.Len(); i++ {
			if arr[i].start <= tmp.end {
				if tmp.end < arr[i].end {
					tmp.end = arr[i].end
				}
			} else {
				fmt.Printf("%s ", tmp)
				tmp.start = arr[i].start
				tmp.end = arr[i].end
			}
		}
		fmt.Printf("%s\n", tmp)
	}
}

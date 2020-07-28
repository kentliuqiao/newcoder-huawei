package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Score int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		d1, _, _ := reader.ReadLine()
		if len(d1) == 0 {
			break
		}
		d2, _, _ := reader.ReadLine()
		if len(d2) == 0 {
			break
		}
		di1, _ := strconv.Atoi(string(d1))
		di2, _ := strconv.Atoi(string(d2))

		var students []Student

		for i := 0; i < di1; i++ {
			tmp, _, _ := reader.ReadLine()
			arr := strings.Split(string(tmp), " ")
			if len(arr) == 2 {
				s1, _ := strconv.Atoi(arr[1])
				students = append(students, Student{
					Name:  arr[0],
					Score: s1,
				})
			}
		}
		// order

		if di2 == 0 {
			for i := 0; i < len(students)-1; i++ {
				for j := 0; j < len(students)-1-i; j++ {
					if students[j].Score < students[j+1].Score {
						students[j], students[j+1] = students[j+1], students[j]
					}
				}
			}
		} else {
			for i := 0; i < len(students)-1; i++ {
				for j := 0; j < len(students)-1-i; j++ {
					if students[j].Score > students[j+1].Score {
						students[j], students[j+1] = students[j+1], students[j]
					}
				}
			}
		}

		// fmt.Println(students)
		for i := range students {
			fmt.Printf("%s %v\n", students[i].Name, students[i].Score)
		}

	}
}

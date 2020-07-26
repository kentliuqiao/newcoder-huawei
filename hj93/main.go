package main

import "fmt"

func main() {
	var total int
	for {
		_, err := fmt.Scan(&total)
		if err != nil {
			break
		}
		tmp := 0
		sums5, sums3 := 0, 0
		others := make([]int, 0)
		for i := 0; i < total; i++ {
			tmp = 0
			fmt.Scan(&tmp)
			if tmp%5 == 0 {
				sums5 += tmp
			} else if tmp%3 == 0 {
				sums3 += tmp
			} else {
				others = append(others, tmp)
			}
		}
		fmt.Println(isExists(sums5, sums3, 0, others))
	}
}

func isExists(sums5, sums3, index int, nums []int) bool {
	if index == len(nums) && sums5 == sums3 {
		return true
	}
	if index == len(nums) && sums3 != sums5 {
		return false
	}
	if index < len(nums) {
		return isExists(sums5+nums[index], sums3, index+1, nums) ||
			isExists(sums5, sums3+nums[index], index+1, nums)
	}
	return false
}

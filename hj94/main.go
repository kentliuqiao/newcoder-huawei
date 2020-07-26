package main

import "fmt"

func main() {
	var candCnt, voteCnt int
	for {
		_, err := fmt.Scan(&candCnt)
		if err != nil {
			break
		}

		cands := make([]string, 0)
		tmp := ""
		for i := 0; i < candCnt; i++ {
			tmp = ""
			fmt.Scan(&tmp)
			cands = append(cands, tmp)
		}

		fmt.Scan(&voteCnt)
		votes := make(map[string]int)
		valid := 0
		for i := 0; i < voteCnt; i++ {
			tmp = ""
			fmt.Scan(&tmp)
			votes[tmp]++
		}
		for i := range cands {
			valid += votes[cands[i]]
			fmt.Println(cands[i], ":", votes[cands[i]])
		}
		fmt.Println("Invalid :", voteCnt-valid)
	}
}

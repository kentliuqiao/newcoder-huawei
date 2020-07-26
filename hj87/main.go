package main

import (
	"fmt"
	"strings"
)

func main() {
	var pass string
	letterlower := "abcdefghijklmnopqrstuvwxyz"
	letterupper := "QWERTYUIOPASDFGHJKLZXCVBNM"
	digits := "0123456789"
	symbols := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	for {
		score := 0
		if _, err := fmt.Scan(&pass); err != nil {
			return

		}

		l := len(pass)
		switch {
		case l <= 4:
			score += 5
		case l <= 7:
			score += 10
		default:
			score += 25
		}

		lower := strings.ContainsAny(pass, letterlower)
		upper := strings.ContainsAny(pass, letterupper)

		if lower && upper {
			score += 20
		} else if lower || upper {
			score += 10
		}

		digitCnt := 0
		for _, b := range pass {
			if strings.Contains(digits, string(b)) {
				digitCnt++
			}
		}
		if digitCnt > 1 {
			score += 20
		} else if digitCnt == 1 {
			score += 10
		}
		symCnt := 0
		for _, b := range pass {
			if strings.Contains(symbols, string(b)) {
				symCnt++
			}
		}
		if symCnt > 1 {
			score += 25

		} else if symCnt == 1 {
			score += 10
		}

		if lower && upper && digitCnt > 0 && symCnt > 0 {
			score += 5
		} else if digitCnt > 0 && symCnt > 0 && (lower || upper) {
			score += 3
		} else if digitCnt > 0 && (lower || upper) {
			score += 2
		}
		switch {
		case score >= 90:
			fmt.Println("VERY_SECURE")
		case score >= 80:
			fmt.Println("SECURE")
		case score >= 70:
			fmt.Println("VERY_STRONG")
		case score >= 60:
			fmt.Println("STRONG")
		case score >= 50:
			fmt.Println("AVERAGE")
		case score >= 25:
			fmt.Println("WEAK")
		default:
			fmt.Println("VERY_WEAK")
		}
	}
}

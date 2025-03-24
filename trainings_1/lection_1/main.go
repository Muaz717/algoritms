package main

import "fmt"

func main() {

	//freqLetter() // O(N^2)
	freqLetterWithMap() // O(N)
}

func freqLetterWithMap() {
	var s string
	fmt.Scan(&s)

	m := make(map[rune]int)

	ansCnt := 0
	ansCh := ""

	for _, ch := range s {
		_, ok := m[ch]
		if !ok {
			m[ch] = 0
		}
		m[ch]++
	}

	for k, v := range m {
		if v > ansCnt {
			ansCh = string(k)
			ansCnt = v
		}
	}
	fmt.Println(m)
	fmt.Println(ansCh)
}

func freqLetter() {
	var s string
	fmt.Scan(&s)

	ansCnt := 0
	ansCh := ""

	for _, ch1 := range s {
		newCnt := 0
		for _, ch2 := range s {
			if ch1 == ch2 {
				newCnt++
			}
		}
		if newCnt > ansCnt {
			ansCh = string(ch1)
			ansCnt = newCnt
		}
	}

	fmt.Println(ansCh)
}

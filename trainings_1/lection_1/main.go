package main

import "fmt"

func main() {

	freqLetter()

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

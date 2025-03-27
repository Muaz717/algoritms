package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list := make([]float64, 0, 2)
	txtNums := make([]string, 0, 2)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		txtNums = append(txtNums, sc.Text())
	}

	strNums := strings.Split(txtNums[1], " ")
	x, _ := strconv.ParseFloat(txtNums[2], 64)

	for _, strNum := range strNums {
		num, _ := strconv.ParseFloat(strNum, 64)
		list = append(list, num)
	}

	ans := list[0]
	diff := math.Abs(x - list[0])

	for _, num := range list {
		newDiff := math.Abs(x - num)

		if newDiff < diff {
			diff = newDiff
			ans = num
		}
	}

	fmt.Println(ans)
}

func nearestNum(list []float64, x float64) {
	ans := list[0]
	diff := math.Abs(x - list[0])

	for _, num := range list {
		newDiff := math.Abs(x - num)

		if newDiff < diff {
			diff = newDiff
			ans = num
		}
	}

	fmt.Println(ans)
}

func seqType(list []int) {
	var countConst, countAsc, countDesc, countWask, countWdesk int

	for i := 0; i < len(list)-1; i++ {

		if list[i+1] == list[i] {
			countConst++
		} else if list[i+1] > list[i] {
			countAsc++
		} else if list[i+1] < list[i] {
			countDesc++
		}

		if list[i+1] >= list[i] {
			countWask++
		}
		if list[i+1] <= list[i] {
			countWdesk++
		}

	}

	size := len(list) - 1

	if countConst == size {
		fmt.Println("CONSTANT")
	} else if countAsc == size {
		fmt.Println("ASCENDING")
	} else if countDesc == size {
		fmt.Println("DESCENDING")
	} else if countWask == size {
		fmt.Println("WEAKLY ASCENDING")
	} else if countWdesk == size {
		fmt.Println("WEAKLY DESCENDING")
	} else {
		fmt.Println("RANDOM")
	}
}

func ascList(seq []int) {

	for i := 0; i < len(seq)-1; i++ {
		if seq[i+1] < seq[i] {
			fmt.Println("NO")
		}
	}

	fmt.Println("YES")
}

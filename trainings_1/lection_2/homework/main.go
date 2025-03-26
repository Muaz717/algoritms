package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list := make([]float64, 0, 2)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		list = append(list, float64(num))
	}
	list = list[0 : len(list)-1]

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

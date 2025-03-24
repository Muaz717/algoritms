package main

import (
	"fmt"
	"strings"
)

//func main() {
//	var a, b, c int
//
//}

func main() {
	num1 := "8(495)430-23-97"
	num2 := "+7-4-9-5-43-023-97"
	num3 := "4-3-0-2-3-9-7"
	num4 := "8-495-430"

	PhoneNumbers(num1, num2, num3, num4)
}

func Сonditioner(a, b int, s string) int {
	switch s {
	case "freeze":
		return min(a, b)
	case "heat":
		return max(a, b)
	case "auto":
		return b
	case "fan":
		return a
	default:
		return a
	}
}

func Triangle(a, b, c int) {

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func PhoneNumbers(num1, num2, num3, num4 string) {
	// Заменить скобки
	// Заменить пробелы
	// Заменить плюсы

	sl := make([]string, 0, 4)
	sl = append(sl, num1, num2, num3, num4)

	for i, str := range sl {
		str = strings.Replace(str, "-", "", -1)
		str = strings.Replace(str, "(", "", 1)
		str = strings.Replace(str, ")", "", 1)
		str = strings.Replace(str, "+", "", 1)
		sl[i] = str
	}

	fmt.Println(sl)
}

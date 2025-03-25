package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	x := rootEquation(1, 2, -3)
	fmt.Println(x)
}

func rootEquation(a, b, c float64) float64 {

	if c < 0 {
		fmt.Println("NO SOLUTION")
		return 0
	} else if a == 0 && b == 0 && c == 0 {
		fmt.Println("MANY SOLUTIONS")
	}

	x := (math.Pow(c, 2) - b) / a

	return x
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

func PhoneNumbers(num1, num2, num3, num4 string) string {

	sl := []string{num1, num2, num3, num4}

	ansSl := make([]string, 0, 3)

	for i, str := range sl {
		str = strings.Replace(str, "-", "", -1)
		str = strings.Replace(str, "(", "", 1)
		str = strings.Replace(str, ")", "", 1)
		str = strings.Replace(str, "+", "", 1)

		if len(str) < 11 {
			str = "8495" + str
		}
		if str[0] == '7' {
			str = strings.Replace(str, "7", "8", 1)
		}

		sl[i] = str
	}

	num1 = sl[0]
	sl = sl[1:]

	for _, str := range sl {

		if num1 == str {
			ansSl = append(ansSl, "YES")
			continue
		} else {
			ansSl = append(ansSl, "NO")
			continue
		}

	}

	return strings.Join(ansSl, ", ")
}

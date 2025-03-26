package main

import "fmt"

func main() {
	ans := pitCraft([]int{3, 1, 2})
	fmt.Println(ans)
}

func twoMax(seq []int) (int, int) {
	mx1 := max(seq[0], seq[1])
	mx2 := min(seq[0], seq[1])

	for _, v := range seq {
		if v > mx1 {
			mx2 = mx1
			mx1 = v
		} else if v > mx2 {
			mx2 = v
		}
	}

	return mx1, mx2
}

func findMinEven(seq []int) int {
	ans := -1

	for _, v := range seq {
		if v%2 == 0 && (ans == -1 || ans > v) {
			ans = v
		}
	}

	return ans
}

func pitCraft(blocs []int) int {
	ans := 0

	mx := blocs[0]
	mxi := 0
	for i, v := range blocs {
		if v > mx {
			mx = v
			mxi = i
		}
	}

	cnt := 1
	for i := 0; i < len(blocs)-1; i++ {

		if i < mxi {
			if blocs[i]-blocs[i+1] == 0 {
				cnt++
			}

			if blocs[i]-blocs[i+1] > 0 {
				ans += (blocs[i] - blocs[i+1]) * cnt
				cnt = 1
			}
		}

		if i > mxi {
			if blocs[i+1]-blocs[i] == 0 {
				cnt++
			}

			if blocs[i+1]-blocs[i] > 0 {
				ans += (blocs[i+1] - blocs[i]) * cnt
				cnt = 1
			}
		}
	}

	return ans
}

package recursion

import "fmt"

func Countdown(num int) {
	fmt.Println(num)
	if num <= 1 {
		return
	}
	num--
	Countdown(num)
}

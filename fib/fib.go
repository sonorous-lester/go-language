package fib

import (
	"fmt"
	"time"
)

func Run(num int) {
	go spinner(100 * time.Millisecond)
	fibN := calculate(num)
	fmt.Printf("\rFibonacci(%d) = %d\n", num, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func calculate(x int) int {
	if x < 2 {
		return x
	}
	return calculate(x-1) + calculate(x-2)
}

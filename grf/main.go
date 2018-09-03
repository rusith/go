package grf

import (
	"fmt"
	"time"
)

func Start() {
	go spinner(100 * time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\f Fibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r);
			time.Sleep(delay);
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x;
	}
	return fib(x-1) + fib(x-2)
}

package udc

import "fmt"

// A simple pipeline to demonstrate the unidirectional channel type

func counter(out chan <- int) { /* Takes an output channel */
	for x := 0; x < 100; x++ { out <- x }
	close(out)
}

func squarer(out chan<- int, in <-chan  int) { /* Takes an output channel and an input channel */
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <- chan int) { /* Takes only an input channel */
	for v := range in {
		fmt.Println(v)
	}
}

func Start() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

package main

import (
	"fmt"
	"runtime"
)

func cpuIntensive(p *int, done chan bool) {
	for i := 0; i <= 100000000; i++ {
		*p = i
	}
	done <- true
}

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool)
	x := 0
	go cpuIntensive(&x, done)

	<-done

	runtime.Gosched()
	fmt.Println("terminou", x)
}

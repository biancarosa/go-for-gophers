package main

import (
	"fmt"
	"runtime"
	"time"
)

func cpuIntensive(p *int) {
	for i := 0; i <= 100000000; i++ {
		*p = i
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	x := 0
	go cpuIntensive(&x)

	time.Sleep(1 * time.Millisecond)

	fmt.Println("terminou", x)
}

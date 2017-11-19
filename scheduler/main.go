package main

import (
	"fmt"
	"runtime"
)

func cpuIntensive(p *int) {
	for i := 0; i <= 100000000; i++ {
		*p = i
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	x := 0
	y := 0
	go cpuIntensive(&x)
	go cpuIntensive(&y)

	runtime.Gosched()
	fmt.Println("terminou", x, y)
}

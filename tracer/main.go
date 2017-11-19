package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func cpuIntensive(p *int) {
	for i := 0; i <= 100000000; i++ {
		*p = i
	}
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	runtime.GOMAXPROCS(1)

	x := 0
	go cpuIntensive(&x)

	runtime.Gosched()
	fmt.Println("terminou", x)
}

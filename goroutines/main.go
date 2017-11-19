package main

import (
	"fmt"
	"runtime"
)

func waitForIt() {
	ch := make(chan bool)
	<-ch
}

func cpuIntensive(p *int) {
	for i := 0; i <= 100000000; i++ {
		go waitForIt()
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	x := 0
	go cpuIntensive(&x)

	runtime.Gosched()
	fmt.Println("terminou", x)
}

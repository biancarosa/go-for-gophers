package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

func waitForIt() {
	ch := make(chan bool)
	<-ch
}

func cpuIntensive(p *int) {
	for i := 0; i <= 100000; i++ {
		go waitForIt()
	}
}
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan os.Signal)
	go func() {
		for range ch {
			pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
		}
	}()
	x := 0
	go cpuIntensive(&x)
	time.Sleep(1 * time.Hour)
	fmt.Println("terminou", x)
}

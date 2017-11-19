package main

import (
	"fmt"
	"runtime"
)

func init() {
	var pcs [1]uintptr
	runtime.Callers(1, pcs[:])
	fn := runtime.FuncForPC(pcs[0])
	fmt.Println(fn.Name())
}

func init() {
	pcs := make([]uintptr, 1)
	runtime.Callers(1, pcs)
	fn := runtime.FuncForPC(pcs[0])
	fmt.Println(fn.Name())
}

func init() {
	fmt.Println("gobr17")
}

func main() {
	fmt.Println("oi")
}

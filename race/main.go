package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := sync.Map{}
	go func() {
		data.Store("name", "bianca")
	}()
	go func() {
		data.Store("surname", "rosa")
	}()

	time.Sleep(1 * time.Second)

	data.Range(func(key interface{}, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

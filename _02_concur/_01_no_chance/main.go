package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64

	go func() {
		for idx := 0; idx < 1_000_000; idx++ {
			atomic.AddInt64(&counter, 1)
			//counter++
		}
	}()

	go func() {
		for idx := 0; idx < 1_000_000; idx++ {
			atomic.AddInt64(&counter, -1)
			//counter--
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", counter)
}

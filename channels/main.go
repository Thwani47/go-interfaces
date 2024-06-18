package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// We use channels to trasnfer data between different go routines
// By default, channels have no space in them. We cannot temporatily store data in them
// A channel is more like a pipe. Data goes in one end and comes out the other end
// To have the data come in one end and go out the other end at the same time, we need to use go routines
// Unbuffered channels have no space in them. They are synchronous
// We can create buffered channels by specifying the buffer size

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n: %d\n", n)
	}
}

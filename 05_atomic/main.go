package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

const numberOfGoRoutines = 500

var counter int32 = 0

func main() {
	newGoRoutine(numberOfGoRoutines)
	wg.Wait()
	fmt.Println("Total go routines", numberOfGoRoutines)
	fmt.Println("counter", counter)
}

func newGoRoutine(n int) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			runtime.Gosched()
			v := atomic.LoadInt32(&counter)
			fmt.Println("Counter", v)
			wg.Done()
		}()
	}

}

// Race condition fixed on exercise number 04 and 05

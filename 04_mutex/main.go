package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

const numberOfGoRoutines = 500

var counter = 0

func main() {
	newGoRoutine(numberOfGoRoutines)
	wg.Wait()
	fmt.Println("Total de go routines", numberOfGoRoutines)
	fmt.Println("Contador", counter)
}

func newGoRoutine(n int) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			mu.Lock()
			c := counter
			runtime.Gosched()
			c++
			counter = c
			mu.Unlock()
			wg.Done()
		}()
	}

}

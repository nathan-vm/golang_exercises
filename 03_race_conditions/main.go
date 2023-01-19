package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const numberOfGoRoutines = 500

var x = 0

func main() {
	newGoRoutine(numberOfGoRoutines)
	wg.Wait()
	fmt.Println("Total de go routines", numberOfGoRoutines)
	fmt.Println("Contador", x)
}

func newGoRoutine(n int) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		c := x
		go func() {
			runtime.Gosched()
			c++
			x = c
			wg.Done()
		}()
	}

}

// Race condition fixed on exercise number 04 and 05

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	newGoRoutine(10)
	wg.Wait()
}

func newGoRoutine(n int) {

	wg.Add(n)

	for i := 0; i < n; i++ {
		x := i
		go func(i int) {
			fmt.Println("I'm go routine number:", i)
			wg.Done()
		}(x)
	}

}

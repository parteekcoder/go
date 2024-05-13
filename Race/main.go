package main

import (
	"fmt"
	"sync"
)

var a []int

// go run --race main.go
func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3)
	go race(1, wg, mut)
	go race(2, wg, mut)
	go race(3, wg, mut)

	wg.Wait()
	fmt.Print(a)
}

func race(i int, wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	a = append(a, i)
	mut.Unlock()
	wg.Done()
}

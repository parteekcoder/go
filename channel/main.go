package main

import (
	"fmt"
	"sync"
)

func main() {
	mych := make(chan int)
	// mych := make(chan int,5) //buffer 5 values

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) { // send only <-chan
		val, isChannelOpen := <-mych
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(mych, wg)
	go func(ch chan int, wg *sync.WaitGroup) { // receive only chan<-
		mych <- 6
		mych <- 5
		close(mych) //closing channel
		wg.Done()
	}(mych, wg)

	wg.Wait()
}

// on listening close channel, it return 0
// sending on channel that is closed will give panic error
// if one listeing or one sending , create deadlock error

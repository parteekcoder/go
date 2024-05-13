package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var i = 0

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Simulate a compute-heavy operation
	sleepDuration := 10 * time.Second

	mu.Lock()
	fmt.Println(i)
	i++
	mu.Unlock()
	time.Sleep(sleepDuration)

	// Send response to the client
	fmt.Fprintln(w, "Request processed!")
}

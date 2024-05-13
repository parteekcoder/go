package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	serverURL   = "http://localhost:8080"
	numRequests = 20
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		fmt.Println(i)
		go func() {
			defer wg.Done()
			sendRequest()
		}()
	}

	wg.Wait()
}

func sendRequest() {
	resp, err := http.Get(serverURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response:", resp.Status)
}

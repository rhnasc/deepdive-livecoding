package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rhnasc/deepdive-livecoding/search"
)

func Google(query string) (results []search.Result) {
	resultCh := make(chan search.Result)

	go func() { resultCh <- First(query, search.Web, search.Web1, search.Web2) }()
	go func() { resultCh <- First(query, search.Image, search.Image1, search.Image2) }()
	go func() { resultCh <- First(query, search.Video, search.Video1, search.Video2) }()

	timeoutCh := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-resultCh:
			results = append(results, result)
		case <-timeoutCh:
			fmt.Println("timed out! :(")
			return
		}
	}

	return
}

func First(query string, replicas ...search.Search) search.Result {
	resultCh := make(chan search.Result)

	for _, replica := range replicas {
		go func() { resultCh <- replica(query) }()
	}

	return <-resultCh
}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())

	results := Google("golang")

	elapsed := time.Since(start)
	fmt.Println("ELAPSED: ", elapsed)
	fmt.Println("RESULTS: ", results)
}

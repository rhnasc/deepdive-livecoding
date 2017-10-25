package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rhnasc/deepdive-livecoding/search"
)

func Google(query string) (results []search.Result) {
	resultCh := make(chan search.Result)

	go func() { resultCh <- search.Web(query) }()
	go func() { resultCh <- search.Image(query) }()
	go func() { resultCh <- search.Video(query) }()

	for i := 0; i < 3; i++ {
		results = append(results, <-resultCh)
	}

	return
}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())

	results := Google("golang")

	elapsed := time.Since(start)
	fmt.Println("ELAPSED: ", elapsed)
	fmt.Println("RESULTS: ", results)
}

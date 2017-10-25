package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rhnasc/deepdive-livecoding/search"
)

func Google(query string) (results []search.Result) {
	results = append(results, search.Web(query))
	results = append(results, search.Image(query))
	results = append(results, search.Video(query))
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

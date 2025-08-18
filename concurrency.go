package main

import "fmt"

//import "time"

// The underlying type of WebsiteChecker is "func(string) bool"
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			fmt.Println("Running goroutine hahahha")
			resultChannel <- result{url, wc(url)}
			//results[url] = wc(url)
		}()
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println("Getting the results from the channel")
		r := <-resultChannel
		results[r.string] = r.bool
	}
	close(resultChannel)
	fmt.Println("To test if the main process waits for the goroutine")

	return results
}


// Run "go test -race" to see the race conditions output

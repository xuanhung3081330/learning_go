package main

import "time"

// The underlying type of WebsiteChecker is "func(string) bool"
type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	time.Sleep(2 * time.Second)

	return results
}


// Run "go test -race" to see the race conditions output

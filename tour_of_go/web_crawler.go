package main

import (
	"fmt"
	"sync"
	"time"
)


type FetchResult struct {
	body string
	urls []string
	err error
}

var cache = make(map[string]*FetchResult)
var mutex sync.Mutex
var threadCount int = 1
var threadMutex sync.Mutex

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func decreaseThreadCount() {
	threadMutex.Lock()
	threadCount--
	threadMutex.Unlock()
}

func increaseThreadCountBy(threads int) {
	threadMutex.Lock()
	threadCount += threads
	threadMutex.Unlock()
}

func getThreadCount() int {
	return threadCount
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer decreaseThreadCount()
	if depth <= 0 {
		return
	}

	var body string
	var urls []string
	var err error

	// Don't fetch the same URL twice.
	mutex.Lock()
	fetchResult, present := cache[url]
	mutex.Unlock()
	if present {
		body, urls, err = fetchResult.body, fetchResult.urls, fetchResult.err
	} else {
		body, urls, err = fetcher.Fetch(url)
		mutex.Lock()
		cache[url] = &FetchResult{body, urls, err}
		mutex.Unlock()
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// Fetch URLs in parallel.
	increaseThreadCountBy(len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	go Crawl("https://golang.org/", 4, fetcher)
	for getThreadCount() > 0 {
		time.Sleep(50 * time.Millisecond)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

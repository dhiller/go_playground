package main

import (
	"fmt"
	"sync"
)

type FetchResult struct {
	body string
	urls []string
	err  error
	url  string
}

func (fetchResult *FetchResult) String() string {
	if fetchResult.err != nil {
		return fmt.Sprint(fetchResult.err)
	}
	return fmt.Sprintf("found: %s %q", fetchResult.url, fetchResult.body)
}

var (
	cache        = make(map[string]*FetchResult)
	mutex        sync.Mutex
	threadCount  = 1
	threadMutex  sync.Mutex
	fetchResults = make(chan *FetchResult)
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (fetchResult *FetchResult)
}

func decreaseThreadCount() {
	threadMutex.Lock()
	defer threadMutex.Unlock()
	threadCount--
	if threadCount <= 0 {
		close(fetchResults)
	}
}

func increaseThreadCountBy(threads int) {
	threadMutex.Lock()
	defer threadMutex.Unlock()
	threadCount += threads
}

func ThreadCount() int {
	return threadCount
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer decreaseThreadCount()
	if depth <= 0 {
		return
	}
	fetchResult := doFetch(url, fetcher)
	fetchResults <- fetchResult
	crawlSubUrls(fetchResult, depth, fetcher)
}

func crawlSubUrls(fetchResult *FetchResult, depth int, fetcher Fetcher) {
	if fetchResult.err != nil {
		return
	}
	increaseThreadCountBy(len(fetchResult.urls))
	for _, u := range fetchResult.urls {
		go Crawl(u, depth-1, fetcher)
	}
}

func doFetch(url string, fetcher Fetcher) *FetchResult {
	mutex.Lock()
	defer mutex.Unlock()
	if _, present := cache[url]; !present {
		cache[url] = fetcher.Fetch(url)
	}
	return cache[url]
}

func main() {
	go Crawl("https://golang.org/", 4, fetcher)
	for fetchResult := range fetchResults {
		fmt.Println(fetchResult)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (fetchResult *FetchResult) {
	if res, ok := f[url]; ok {
		return &FetchResult{body: res.body, urls: res.urls, url: url}
	}
	return &FetchResult{err: fmt.Errorf("not found: %s", url), url: url}
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

package main

import (
	"fmt"
	"sync"
	//"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var (
	mutex sync.Mutex
	waitG sync.WaitGroup
	cache = make(map[string]bool)
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer waitG.Done()
	//fmt.Println("----------------START-----------------")
	//for u, b := range cache {
	//	fmt.Printf("Cache: %s, %b\n", u, b)
	//}
	//fmt.Println("-----------------END------------------")
	if cache[url] {
		mutex.Lock()
		//fmt.Println("Already crawled: ", url)
		mutex.Unlock()
		return
	}
	//fmt.Println("Crawling: ", url)
	mutex.Lock()
	cache[url] = true
	mutex.Unlock()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		waitG.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	//waitG.Done()
	return
}

func main() {
	waitG.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	//time.Sleep(time.Second * 5)
	waitG.Wait()
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

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0x0)
// 	/usr/local/go-faketime/src/runtime/sema.go:56 +0x25
// sync.(*WaitGroup).Wait(0x49a702)
// 	/usr/local/go-faketime/src/sync/waitgroup.go:130 +0x71
// main.main()
// 	/tmp/sandbox2165778371/prog.go:58 +0x55

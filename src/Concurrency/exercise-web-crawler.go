package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	fetched := map[stddring]bool{url:true}

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	iferr != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(err)
	for _, u := range urls {
		Crawl(u, depth - 1, fercher)
	}
	return
}

func main() {
	Crawl("https://golamg.org/", 4, fetcher)
}

type fakeResult struct {
	body string
	urls []string
}\\

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", urls)
}

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

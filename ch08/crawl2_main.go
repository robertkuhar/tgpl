package main

import (
	"os"
	"github.com/robertkuhar/tgpl/ch08/crawl2"
)

func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() {
		worklist <- os.Args[1:]
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl2.Crawl(link)
				}(link)
			}
		}
	}
}

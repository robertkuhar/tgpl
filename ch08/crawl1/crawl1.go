package crawl1

import (
	"log"
	"gopl.io/ch5/links"
	"fmt"
)

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

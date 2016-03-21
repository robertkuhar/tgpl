package crawl2

import (
	"log"
	"gopl.io/ch5/links"
	"fmt"
)

var tokens = make(chan struct{}, 20)

func Crawl(url string) []string {
	fmt.Println(url)
	// TODO: BobK I still don't get the struct{}{} syntax
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

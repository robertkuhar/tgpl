package main

import (
	"flag"
	"time"
	"fmt"
)

/*
This is going to take some getting used to.  Why is the return value an address and not just the duration itself?
TODO:  How can I get IDEA to format this comment properly with line wrap at the margin and all.
 */
var period = flag.Duration("period", 1 * time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("At %s, Sleeping for %v...\n", time.Now().UTC(), *period)
	time.Sleep(*period)
	fmt.Printf("At %s, Awakened\n", time.Now().UTC())

	// Boy, this is bizarre.  time.Format works from some proto-typie thing based in 2006!?!
	t := time.Now()
	fmt.Printf("UTC\t\t\t\t%s\n", t.UTC())
	robertDate := "2006-01-02T15:04:05.000Z"
	fmt.Printf("UnixDate:\t\t%s\n", t.Format(time.UnixDate))
	fmt.Printf("my guess:\t\t%s\n", t.UTC().Format(robertDate))
}
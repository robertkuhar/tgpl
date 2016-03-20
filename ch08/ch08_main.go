package main

import (
	"fmt"
	"strings"
)

func generator(out chan <- string) {
	out <- "Bob"
	out <- "Robert"
	out <- "Dread Pirate Roberts"
}

func uni_tagger(out chan <- string, in <-chan string) {
	for s := range in {
		out <- fmt.Sprintf("tagged %s", s)
	}
}

func re_tagger(inout chan string) {
	for s := range inout {
		if !strings.Contains(s, "re_tagger") {
			inout <- fmt.Sprintf("re-tagged %s", s)
		}
	}
}

func printer(in <-chan string) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	inputs := make(chan string)
	outputs := make(chan string)

	go generator(inputs)
	go uni_tagger(outputs, inputs)
	go re_tagger(inputs)
	printer(outputs)
}

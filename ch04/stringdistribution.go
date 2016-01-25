package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func main() {
	candidates := []string{"One", "Two", "Two" }
	Add(candidates)
	Add(candidates)
	others := []string{"one","two"}
	Add(others)
	fmt.Printf("Count: %v %v\n", candidates, Count(candidates))
	fmt.Printf("Count: %v %v\n", others, Count(others))
}

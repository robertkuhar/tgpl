package main
import (
	"fmt"
	"sort"
)

type name_and_age struct {
	name string
	age  int
}

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages["bob"] = 51
	fmt.Printf("bob %v, tom %v, alice %v\n", ages["bob"], ages["tom"], ages["alice"])
	ages["tom"] += 1
	fmt.Printf("bob %v, tom %v, alice %v\n", ages["bob"], ages["tom"], ages["alice"])
	fmt.Printf("sorted %v\n",sortMap(ages))
}

func sortMap(src map[string]int) []name_and_age {
	var names [] string
	var sorted []name_and_age
	for name := range src {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, src[name])
		var name_and_age_instance name_and_age
		name_and_age_instance.name = name
		name_and_age_instance.age = src[name]
		sorted = append(sorted, name_and_age_instance)
	}
	return sorted
}


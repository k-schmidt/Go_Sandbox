package main

import "fmt"

// main ...
func main() {
	a := [...]int{1, 2, 3, 4, 5}
	ages := map[string]int{
		"alice": 31,
		"bob":   34,
	}

	fmt.Println(a)
	fmt.Println(ages)
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

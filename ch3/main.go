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

	// key lookups on map never produce errors
	fmt.Println(ages["carol"]) // == 0 DNE
	if _, ok := ages["carol"]; !ok {
		fmt.Printf("Carol DNE!!!\nok == %v", ok)
	}
}

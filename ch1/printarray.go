package main

import (
	"fmt"
	"os"
)

var names = [2]string{"Kyle", "Schmidt"}
var boiling float64 = 100

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var v, w int
	fmt.Println(&v == &v, &v == &w, &v == nil)

	s := 1
	incr(&s)
	fmt.Println(incr(&s))

	for _, f := range names {
		fmt.Fprintf(os.Stdout, "%f\n", boiling)
		fmt.Fprintf(os.Stdout, "%s\n", f)
	}
}

// incr increments what p points to; does not change p
func incr(p *int) int {
	*p++
	return *p
}

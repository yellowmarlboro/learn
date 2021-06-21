package main

import "fmt"

func main() {
	p := new([]int) //p == nil; with len and cap 0
	fmt.Println(p)

	v := make([]int, 10, 50) // v is initialed with len 10, cap 50
	fmt.Println(v)

	/*********Output****************
	    &[]
	    [0 0 0 0 0 0 0 0 0 0]
	*********************************/

	(*p)[0] = 18        // panic: runtime error: index out of range
	// because p is a nil pointer, with len and cap 0
	v[1] = 18           // ok

}

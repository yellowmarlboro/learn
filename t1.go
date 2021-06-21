package main

import "fmt"

func main()  {
	s1 := [...]int {1, 2, 3}
	fmt.Printf("v1 type:%T\n", s1)

	s2 := [3]int {1, 2, 3}
	fmt.Printf("v1 type:%T\n", s2)

	s3 := []int {1, 2, 3}
	fmt.Printf("v1 type:%T\n", s3)
}

package inside

import "fmt"

func modify(slice []int) {
	slice = append(slice, 100)
}

func sliceL() {
	s := make([]int, 3, 10)
	s1 := make([]int, 3, 3)
	fmt.Println(s)
	fmt.Println(s1)
	modify(s)
	modify(s1)
	fmt.Println(s)
	fmt.Println(s1)
}
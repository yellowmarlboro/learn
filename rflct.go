package main

import (
	"encoding/json"
	"fmt"
)

func ReflectT1()  {
	a := A{1, 2}
	b := B{10, 20}
	bA, _ := json.Marshal(a)
	bB, _ := json.Marshal(b)

	fmt.Println(bA, bB)
}
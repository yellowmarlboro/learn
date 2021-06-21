package main

import "fmt"

type VerTex struct {
	Lat, Long float64
}

var m map[string] VerTex

func main()  {
	// TODO 为什么要再make一次啊
	m = make(map[string]VerTex)
	m["test"] = VerTex{
		40, 40.1,
	}
	fmt.Println(m)
}
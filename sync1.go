package main

import (
	"fmt"
	"sync"
	"time"
)

var wc sync.WaitGroup

func main()  {
	wc.Add(1)
	go func() {
		defer wc.Done()
		fmt.Println("4")
		fmt.Println("5")
		fmt.Println("6")
		fmt.Println("7")
	}()
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")
	wc.Wait()
	go func() {
		fmt.Println("14")
		fmt.Println("15")
		fmt.Println("16")
		fmt.Println("17")
	}()
	fmt.Println("18")
	fmt.Println("19")
	fmt.Println("20")
	time.Sleep(2)
}

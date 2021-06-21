package main

import (
	"fmt"
	"sync"
	"time"
)

var wC1 sync.WaitGroup

// 无缓冲channel无论发送操作还是接受操作一开始就是阻塞的，只有配对的操作出现才会开始执行
// 有缓冲这会在 已满/已空情况下的发送/接受操作会阻塞

func main() {
	wC1.Add(6)
	chNoBuffer := make(chan int)
	chBuffer1 := make(chan int, 3)
	go func() {
		defer wC1.Done()
		fmt.Println("start consumer nobuffer chan")
		fmt.Printf("nobuffer chan is %d \n", <-chNoBuffer)
	}()
	go func() {
		defer wC1.Done()
		fmt.Println("start product nobuffer chan")
		time.Sleep(3 * time.Second)
		chNoBuffer <- 123
		fmt.Println("product nobuffer chan done")
	}()
	go func() {
		defer wC1.Done()
		fmt.Println("start product nobuffer chan 2")
		time.Sleep(6 * time.Second)
		chNoBuffer <- 456
		fmt.Println("product nobuffer chan done 2")
	}()
	go func() {
		defer wC1.Done()
		fmt.Println("start consumer nobuffer chan 2")
		fmt.Printf("nobuffer chan is %d 2\n", <-chNoBuffer)
	}()
	go func() {
		defer wC1.Done()
		fmt.Println("start consumer chBuffer1 chan")
		fmt.Printf("chBuffer1 chan is %d \n", <-chBuffer1)
	}()
	go func() {
		defer wC1.Done()
		fmt.Println("start product chBuffer1 chan")
		chBuffer1 <- 789
		fmt.Println("product chBuffer1 chan done")
	}()
	wC1.Wait()
	fmt.Println("done")
}

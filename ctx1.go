package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var(
	ch1 = make(chan int)
	ch2 = make(chan int)
	ch3 = make(chan int)
	wg = sync.WaitGroup{}
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 4*time.Second)

	go Ctx1(ctx)
	select {
	case <- ctx.Done():
		fmt.Println("main Failed", ctx.Err())
	case ret := <- ch1:
		fmt.Printf("result is %d", ret)
	}
	time.Sleep(10*time.Second)
}

func Ctx1(ctx context.Context){
	wg.Add(4)
	ctx, cancleFunc := context.WithCancel(ctx)
	go Ctx2(ctx, cancleFunc)
	go Ctx3(ctx, cancleFunc)

	var lData []int
	go func() {
		select {
		case s := <- ctx.Done():
			fmt.Println("get data3 is err, ", ctx.Err(), s)
			cancleFunc()
		case data := <- ch3:
			fmt.Println("get data3 Done", data)
			lData = append(lData, data)
		}
		wg.Done()
	}()

	go func() {
		select {
		case s := <- ctx.Done():
			fmt.Println("get data2 is err, ", ctx.Err(), s)
			cancleFunc()
		case data := <- ch2:
			fmt.Println("get data2 Done", data)
			lData = append(lData, data)
		}
		wg.Done()
	}()

	wg.Wait()

	Sum := sum(lData)
	fmt.Println(Sum, "done logic1.")
	ch1 <- Sum
}

func Ctx2(ctx context.Context, f context.CancelFunc) {
	fmt.Println("in Ctx2.")
	select {
	case <- ctx.Done():
		//f()
		fmt.Println("get cancle single 2")
		return
	case <- time.After(6*time.Second):
		ch2 <- 10
	}
	fmt.Println("done Ctx2.")
	wg.Done()
}

func Ctx3(ctx context.Context, f context.CancelFunc) {
	fmt.Println("in Ctx3.")
	select {
	case <- ctx.Done():
		//f()
		fmt.Println("get cancle single 3")
		return
	case <- time.After(4*time.Second):
		ch3 <- 12
	}
	fmt.Println("done Ctx3.")
	wg.Done()
}

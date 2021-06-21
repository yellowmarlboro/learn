package main
//
//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//func main() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	for i := 0; i <= 100; i++ {
//		wg.Add(1)
//		go add(i)
//
//	}
//	wg.Wait()
//}
//
//var s = 1000
////var lock sync.Mutex
//var wg = &sync.WaitGroup{}
//
//func add(count int) {
//	//lock.Lock()
//	fmt.Printf("加锁----第%d个携程\n", count)
//	for i := 0; i < 4; i++ {
//		s++
//		fmt.Printf("j %d gorount %d \n", s, count)
//	}
//	fmt.Printf("解锁----第%d个携程\n", count)
//	wg.Done()
//	//defer lock.Unlock()
//}
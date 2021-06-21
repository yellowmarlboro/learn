package main

import "fmt"

func main()  {
	s := []int {1, 2, 3, 4, 5, 6}
	printSlice(s)

	// 改变底层数组的值
	newS := s[:3]
	newS[2] = 100
	printSlice(newS)

	// 定义切片时，控制容量（容量从底层数组下标0开始，比如4代表【1,2,3,4】
	// 内置函数 append() 在操作切片时会首先使用可用容量。一旦没有可用容量，就会分配一个新的底层数组。
	// 这导致很容易忘记切片间正在共享同一个底层数组。一旦发生这种情况，对切片进行修改，很可能会导致随机且奇怪的问题
	// 这种问题一般都很难调查。如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个 append 操作创建新的底层数组
	// 与原有的底层数组分离。这样就可以安全地进行后续的修改操作
	newS2 := s[3:4:4]
	newS2 = append(newS2, s...)
	printSlice(newS2)

	s = s[:2]
	printSlice(s)

	s = s[1:]
	printSlice(s)

	for idx, val := range newS2{
		fmt.Println(idx, val)
	}

	num1 := []int{10, 20, 30}
	num2 := make([]int, 5)
	// 切片返回拷贝成功的个数
	count := copy(num2, num1)
	fmt.Println(count)
	fmt.Println(num2)

	n := 10
	print(n)

}

func printSlice(s[]int)  {
	fmt.Printf("len=%d cap=%d v is %v \n", len(s), cap(s), s)
}

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)
//
//func modifyStr(a []int, a1[5]int, b *string, args ...string)  {
//	a[0] = 100
//	a1[0] = 200
//	*b += "A"
//	for _, val := range args{
//		val = val + "B"
//	}
//}
//
type Info struct {
	Sex int
	Age int
}
type Student struct {
	Name string  `json:"first_name" tt:"test_tag"`
	Info
	Class string
}

func (i Info)printSex() {
	if i.Sex == 1{
		fmt.Println("男")
	} else {
		fmt.Println("女")
	}
}

func (s *Student) upLevel () {
	s.Info.Age += 1
	s.Class += "+"
}


func main() {
	man1 := Student{Name: "bob", Class: "S"}
	t := reflect.TypeOf(man1)
	f, _ := t.FieldByName("Name")
	fmt.Println(f.Tag) // json:"first_name" tt:"test_tag"

	man1.printSex() // 女
	man1.upLevel()
	fmt.Println(man1) // {bob {0 1} S+}

	data, err := json.Marshal(man1)
	if err != nil{
		fmt.Println("error")
	}
	fmt.Println(string(data)) // {"first_name":"bob","Sex":0,"Age":0,"Class":"S"}

	//s := [] int {1, 2, 3, 4}
	//s1 := s
	//s2 := make([]int, 4, 8)
	//
	//copy(s2, s1)
	//s1[0] = 100
	//fmt.Println(s, s1, s2)
	//// [100 2 3 4] [100 2 3 4] [1 2 3 4]
	//fmt.Printf("s len is %d cap is %d \n", len(s), cap(s))
	//// s len is 4 cap is 4
	//fmt.Printf("s2 len is %d cap is %d", len(s2), cap(s2))
	// s2 len is 4 cap is 8
	//n := new([]int)
	//fmt.Println(n)
	//m := make([]int, 4, 6)
	//fmt.Printf("m is %v, len is %d, cap is %d", m, len(m), cap(m))
	//var arr1 [5]int
	//arr2 := [5]int{10,20,30,40,50}
	//arr3 := [...]int{10,20,30,40,50}
	//arr4 := [5]int{1:90,3:80}
	//fmt.Println(arr1, arr2, arr3, arr4)
	// [0 0 0 0 0] [10 20 30 40 50] [10 20 30 40 50] [0 90 0 80 0]
	//a := []int {1,2,3,4,5}
	//a1 := [5]int {1,2,3,4,5}
	//b := "something"
	//c := "type"
	//d := "func"
	//modifyStr(a, a1, &b, c, d)
	//fmt.Println(a, a1, b, c, d)
	// [100 2 3 4 5] [1 2 3 4 5] somethingA type func

	//var b byte = 101 // e的ascii码
	//var b1 byte = 'e'
	//var r rune = 22909
	//fmt.Println(string(b)) // e
	//fmt.Println(b) // 101
	//fmt.Println(string(b1)) // e
	//fmt.Println(b1) // 101
	//fmt.Println(string(r)) // 好
	//fmt.Println(r) // 22909
	//
	//var s string = "你好 world"
	//fmt.Println(s)
	//fmt.Println(len(s)) // 12
	//fmt.Println(s[0]) // 228
	//fmt.Println(s[7]) // 119
	//fmt.Println([]rune(s))
	//// [20320 22909 32 119 111 114 108 100]
	//fmt.Println([]byte(s))
	//// [228 189 160 229 165 189 32 119 111 114 108 100]
	//
	//for _, val := range(s){
	//	fmt.Println(string(val))
	//}

	//s0 := "something"
	//s1 := "something"
	//s2 := "something"[7:]
	//fmt.Println(&s0, &s1, &s2)
	//fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s0)).Data)
	//fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Data)
	//fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Data)

	//i := 12
	//switch {
	//case i % 7 == 0:
	//	fmt.Println("7")
	//case i % 6 == 0:
	//	fmt.Println("6")
	//	fallthrough
	//case i % 3 == 0:
	//	fmt.Println("3")
	//default:
	//	fmt.Println("default")
	//s := "something"
	//for _, val := range s {
	//	val += 1
	//}
	//fmt.Println(s)

}


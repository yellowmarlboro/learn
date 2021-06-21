package stream

import (
	"fmt"
)

type DataSource interface {

}

type DataChan chan interface{}

type DataProc struct {
	Entity UuserEntity
	LeftNodes []DataProc
	RightNodes []DataProc
	RightChan []chan UuserEntity
	OutputVars []string
}

type B struct {
	Class string
}

type AAA struct {
	Name string
	Age int
	B  B
}

func main() {
	s := AAA{
		Name: "Sw",
		Age:  18,
		B:    B{"test"},
	}
	fmt.Println(fmt.Sprintf("err is %d+", s))
}
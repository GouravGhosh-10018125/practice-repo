package package1

import "fmt"

type MyStruct struct {
	x int
}

func NewMyStruct(x int) *MyStruct {
	return &MyStruct{x: x}
}

func (m *MyStruct) MyMethod() {
	fmt.Println("hi from myMethod, but you can't see me cuz i'm unexported lol")
}

var UNCLE = 89

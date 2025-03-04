package main

import (
	"fmt"

	"go-practice/package1"
)

func main() {
	fmt.Println("Hello, World!")
	myStruct := package1.NewMyStruct(4)
	(*myStruct).MyMethod()
	fmt.Println(*myStruct)
	var UNCLE = package1.UNCLE
	fmt.Println(UNCLE)
}

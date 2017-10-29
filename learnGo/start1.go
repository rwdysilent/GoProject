package main

import (
	"fmt"
	"math"
)

func foo() int {
	x := 3
	y := 2
	return x + y
}

func foo1(x, y int) int {
	return x + y
}

//多指返回
func foo2(x, y int) (a, b int) {
	a = x
	b = y
	return
}

//结构体
type Vertex struct {
	X int
	Y int
}

type Vertex1 struct {
	Z int
}

func main() {
	//
	fmt.Println(foo())
	//
	x := 3
	y := 2
	fmt.Println(foo1(x, y))
	//
	fmt.Println(foo2(x, y))

	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	fmt.Println(Vertex1{4})

	fmt.Println(math.Sqrt(25))
}

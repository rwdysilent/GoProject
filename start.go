package main

import (
	"fmt"
	"math/rand"
	"math"
)

//var c, python, java bool

func add(x, y int) int {
	return x + y
}

//定义变量


//变量类型转换
var (
	i int = 42
	f float64 = float64(i)
	u uint = uint(f)
)

//常量定义
const name  = "world"

func main() {
	//:= 结构不能使用在函数外
	c, python, java := true, false, "no!"
	fmt.Println(math.Pi)
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(42, 13))
	fmt.Println(i, c, python, java, u, name)
}

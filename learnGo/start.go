package main

import (
	"fmt"
	"math/rand"
	"math"
	"runtime"
	"time"
)

//var c, python, java bool

func add(x, y int) int {
	return x + y
}

//定义变量
var a, b int = 3, 4
//var c, python, java bool

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

	//switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	//switch 顺序执行
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//defer 栈
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	fmt.Println(i, c, python, java, u, name, a, b)
}

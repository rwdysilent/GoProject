package main

import (
	"fmt"
	"math/rand"
	"math"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(math.Pi)
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(42, 13))
}

package main

import (
	"fmt"
)

func main() {
	b := []int{3, 4}
	var x int
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
}
func nextInt(ints []int, i int) (int, int) {

}

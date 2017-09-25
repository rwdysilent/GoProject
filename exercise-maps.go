package main

import (
	"strings"
	"fmt"
)

var s string = "I love wxx , I love wpf"

func LearnStrings(s, w string) bool{
	return strings.Contains(s, w)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	word := strings.Fields(s)
	for _,value := range word{
		m[value]++
	}
	fmt.Println(m)
	return m
}

func main()  {
	a := strings.Fields(s)
	for i, j := range a{
		fmt.Println(i,j)
	}
	WordCount(s)

	fmt.Println(LearnStrings(s,"wxx"))

}

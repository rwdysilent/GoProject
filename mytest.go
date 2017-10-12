package main

import "fmt"

func Append(slice, data[]byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {
		// 重新分配
		// 为了后面的增长，需分配两份。
		newSlice := make([]byte, (l + len(data)) * 2)
		// copy 函数是预声明的，且可用于任何切片类型。
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l + len(data)]
	for i, c := range data {
		slice[l + i] = c
	}
	return slice
}

func main_back() {
	s := make([]byte, 0, 5)
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 2, 3, 4, 5, 6)
	d := []byte{1, 2, 3, 4, 5, 6}
	fmt.Println(len(s), cap(s))
	var mews[]byte
	mews = Append(s, d)
	fmt.Println(mews)

	t := s[0:1]
	fmt.Println(t, s)
}

func main() {
	type LinesOfText [][]byte
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	fmt.Println(text)
}
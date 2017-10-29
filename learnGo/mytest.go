package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
)

//func Append(slice, data []byte) []byte {
//	l := len(slice)
//	if l+len(data) > cap(slice) {
//		// 重新分配
//		// 为了后面的增长，需分配两份。
//		newSlice := make([]byte, (l+len(data))*2)
//		// copy 函数是预声明的，且可用于任何切片类型。
//		copy(newSlice, slice)
//		slice = newSlice
//	}
//	slice = slice[0:l+len(data)]
//	for i, c := range data {
//		slice[l+i] = c
//	}
//	return slice
//}
//
//func main_back() {
//	s := make([]byte, 0, 5)
//	fmt.Println(len(s), cap(s))
//	s = append(s, 1, 2, 3, 4, 5, 6)
//	d := []byte{1, 2, 3, 4, 5, 6}
//	fmt.Println(len(s), cap(s))
//	var mews []byte
//	mews = Append(s, d)
//	fmt.Println(mews)
//
//	t := s[0:1]
//	fmt.Println(t, s)
//}

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func offset(tz string) int {
	seconds, ok := timeZone[tz]
	fmt.Println(seconds, ok)
	if seconds, ok := timeZone[tz]; ok {
		fmt.Println("OK: %s", ok)
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}

type Stereotype int

//func main() {
//	x := []int{1, 2, 3}
//	y := []int{4, 5}
//	x = append(x, y...)
//	fmt.Println(x)
//
//	const (
//		TypicalNoob           Stereotype = iota // 0
//		TypicalHipster                          // 1
//		TypicalUnixWizard                       // 2
//		TypicalStartupFounder                   // 3
//	)
//	fmt.Println(TypicalNoob, TypicalHipster, TypicalUnixWizard, TypicalStartupFounder)
//
//	user := ""
//	fmt.Println(user)
//
//	m := MyFunc{0,1,2,3,4}
//	fmt.Println(m.abs())
//}

func init() {
	fmt.Println("I am init")
}

type MyFunc []int

func (m *MyFunc) abs() (slice []int) {
	p := *m
	n := p[0:3]
	*m = n
	return *m
}

type Sequence []int

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	sort.Sort(s)
	//str := "["
	//for i, elem := range s {
	//    if i > 0 {
	//        str += " "
	//    }
	//    str += fmt.Sprint(elem)
	//}
	//m_str := str + "]"
	//fmt.Println(reflect.TypeOf(s))
	return fmt.Sprint([]int(s))
}

type myInteger int

func (p myInteger) get() int {
	return int(p)
} // Conversion required.

func f1(i myInteger) myInteger {
	return i
}

func f2(i int) float64 {
	return float64(i)
}

var v myInteger

type myType struct{ i int }

func (p *myType) get() int { return p.i }

func (p *myType) set(i int) { p.i = i + 3 }

type myInterface interface {
	get() int
	set(i int)
}

func getAndSet(x myInterface) int {
	//x.set(3)
	return x.get()
}

func ff() int {
	var p myType
	//p := myType{5}
	a := getAndSet(&p)
	return a
}

type myChildType struct {
	myType
	j int
}

func (p *myChildType) get() int {
	n := p.j + 1
	p.myType.set(n)
	fmt.Printf("n is %d\n", n)
	return p.myType.get()
}

func ff2() (int, int) {
	//var p myChildType
	//q := myType{1}
	p := myChildType{myType{1}, 2}
	a := getAndSet(&p)
	return a, p.j
}

//type content interface {}

func echoString() {
	var str interface{}
	//str := 1
	//result, _ := content.(string)  //通过断言实现类型转换
	if i, ok := str.(string); ok {
		fmt.Println(i)
	} else {
		fmt.Printf("Err: %T-->%v\n", str, str)
	}
}

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	fmt.Println(a)
	fmt.Println(s)
	return
}

type Work string

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

func do(work *Work) {
	panic(fmt.Sprintf("飘了: %s", work))
}

func main() {
	//a := regexp.MustCompile("a")
	//fmt.Println(a.Split("banana", -1))
	//fmt.Println(a.Split("banana", 0))
	//fmt.Println(a.Split("banana", 1))
	//fmt.Println(a.Split("banana", 2))
	//zp := regexp.MustCompile("z+")
	//fmt.Println(zp.Split("pizza", -1))
	//fmt.Println(zp.Split("pizza", 0))
	//fmt.Println(zp.Split("pizza", 1))
	//fmt.Println(zp.Split("pizza", 2))

	str := "/playlist?id=965267769"
	s := regexp.MustCompile("\\?id=")
	ans := s.Split(str, 2)
	m_sli := make([]string, 0)
	m_sli = append(m_sli, ans[1])
	fmt.Println(ans[1], m_sli, len(m_sli))
}

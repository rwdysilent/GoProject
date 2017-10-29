package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}


type Err struct {}
func (_ *Err) Error() string {
        return "To err is human"
}

func NoErr(ok bool) error {
        if !ok {
                return &Err{}
        }
        return nil
}

func ToErr(ok bool) error {
        var e *Err = nil
        if ok {
                e = &Err{}
        }
        return e
}

func main() {
	//if err := run(); err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println(NoErr(true))
        fmt.Println(NoErr(false))

	fmt.Println(ToErr(true))  //false
    	fmt.Println(ToErr(false))  //false
}

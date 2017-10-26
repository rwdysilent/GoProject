//refer: https://gobyexample.com/json
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"os"
)

type Response1 struct {
    Page   int
    Fruits []string
}
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
	// First we'll look at encoding basic data types to
    // JSON strings. Here are some examples for atomic
    // values.
	bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

	// And here are some for slices and maps, which encode
    // to JSON arrays and objects as you'd expect.
	fmt.Println("------------------")
	slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB), reflect.TypeOf(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB), reflect.TypeOf(mapB))

	// The JSON package can automatically encode your
    // custom data types. It will only include exported
    // fields in the encoded output and will by default
    // use those names as the JSON keys.
	fmt.Println("res1B res2B------------------")
	res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
	// You can use tags on struct field declarations
    // to customize the encoded JSON key names. Check the
    // definition of `Response2` above to see an example
    // of such tags.
	res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

	// Now let's look at decoding JSON data into Go
    // values. Here's an example for a generic data
    // structure.
	fmt.Println("------------------")
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	// Here's the actual decoding, and a check for
    // associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

	// In order to use the values in the decoded map,
    // we'll need to cast them to their appropriate type.
    // For example here we cast the value in `num` to
    // the expected `float64` type.
    num := dat["num"].(float64)
    fmt.Println(num)

	// Accessing nested data requires a series of
    // casts.
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

	// We can also decode JSON into custom data types.
    // This has the advantages of adding additional
    // type-safety to our programs and eliminating the
    // need for type assertions when accessing the decoded
    // data.
	fmt.Println("------------------")
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits)

	// In the examples above we always used bytes and
    // strings as intermediates between the data and
    // JSON representation on standard out. We can also
    // stream JSON encodings directly to `os.Writer`s like
    // `os.Stdout` or even HTTP response bodies.
	fmt.Println("------------------")
	enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
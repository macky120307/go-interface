package main

import "fmt"

type Func func() string
func (f Func) String() string { return f() }
func main() {
	var s fmt.Stringer = Func(func() string {
		return "hi"
	})
	fmt.Println(s)

	var i interface{}
	i = nil
	switch v := i.(type) {
	case int:
		fmt.Println(v*2)
	case string:
		fmt.Println(v+"hoge")
	default:
		fmt.Println("default")
	}
}


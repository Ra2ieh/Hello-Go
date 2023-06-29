package main

import "fmt"

func main() {
	const a = 42        //untyped constant
	const b bool = true //typed constant
	const c = "hello"

	const (
		d = 43
		f = 50
		s = "hi"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(s)
	// has error :f++
}

package main

import "fmt"

type razi int

var r razi

func main() {
	fmt.Println(r)
	fmt.Printf("%T\n", r)
	r = 42
	println(r)
}

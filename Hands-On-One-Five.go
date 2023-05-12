package main

import "fmt"

type razi1 int

var w razi1

func main() {
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	w = 42
	fmt.Println(w)
	c := int(w)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}

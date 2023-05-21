package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	b = 43
	var a int
	//a=b error
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//conversion
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T", a)
}

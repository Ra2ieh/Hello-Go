package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[2] = 4
	fmt.Println(x)
	fmt.Println(len(x))
	a := [4]int{10, 20, 30, 40}
	fmt.Println(a)
	var b [5]int
	fmt.Println(b)
	// has error : a=b
}

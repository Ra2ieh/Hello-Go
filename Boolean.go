package main

import "fmt"

func main() {
	var p bool
	fmt.Println(p)
	p = true
	fmt.Println(p)
	a := 7
	c := 42
	fmt.Println(a == c)
	fmt.Println(a != c)
	fmt.Println(a >= c)
	fmt.Println(a <= c)
	fmt.Println(a == c/6)

}

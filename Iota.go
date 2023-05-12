package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
	)

	const (
		f = iota
		g
		h
		o
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(o)
}

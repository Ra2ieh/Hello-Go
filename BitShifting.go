package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("x in decimal is : %d\n", x)
	fmt.Printf("x in binary is : %b\n", x)
	y := x << 1
	fmt.Printf("y in decimal is : %d\n", y)
	fmt.Printf("y in binary is : %b\n", y)

	z := x >> 1
	fmt.Printf("z in decimal is : %d\n", z)
	fmt.Printf("z in binary is : %b\n", z)

	w := y << 1
	fmt.Printf("w in decimal is : %d\n", w)
	fmt.Printf("w in binary is : %b\n", w)
	fmt.Println("--------------------------------------------------")
	//kb := 1024
	//mb := kb * 1024
	//gb := mb * 1024

	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)
	fmt.Printf("kb in decimal is : %d\n", kb)
	fmt.Printf("mb in decimal is : %d\n", mb)
	fmt.Printf("gb in decimal is : %d\n", gb)

	fmt.Printf("kb in binary is : %b\n", kb)
	fmt.Printf("mb in binary is : %b\n", mb)
	fmt.Printf("gb in binary is : %b\n", gb)

}

package main

import "fmt"

func main() {
	a := 10
	b := 100
	s := "hello"
	fmt.Printf("the value of a is %v and the address is %v\n", a, &a)
	fmt.Printf("the value of b is %v and the address is %v\n", b, &b)
	fmt.Printf("the value of s is %v and the address is %v\n", s, &s)
	increment(a)
	fmt.Printf("the value of a is %v and the address is %v\n", a, &a)
	b = increase(a)
	fmt.Printf("the value of b is %v and the address is %v\n", b, &b)
}
func increment(v int) {
	fmt.Printf("the value of v before operation is %v and the address is %v\n", v, &v)
	v++
	fmt.Printf("the value of v after operation is %v and the address is %v\n", v, &v)
}
func increase(i int) int {
	fmt.Printf("the value of i before operation is %v and the address is %v\n", i, &i)
	i++
	fmt.Printf("the value of i after operation is %v and the address is %v\n", i, &i)
	return i
}

package main

import "fmt"

func main() {
	fmt.Println("before foo function")
	foo()
	println("after foo function")
}
func foo() {
	println("in the foo function")
}

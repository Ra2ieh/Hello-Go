package main

import "fmt"

func main() {
	n := 1
	//declare an anonymous function and call it
	func() {
		fmt.Println("n is : ", n)
	}()
	n++
	//declare an anonymous function and assign it to a variable
	f := func() {
		fmt.Println("n is : ", n)
	}
	f()
	n++
	//defer the call of anonymous function till after main returns
	defer func() {
		fmt.Println("Defer func1 is : ", n)
	}()
	f()
	n++
	defer func() {
		fmt.Println("Defer func2 is : ", n)
	}()
}

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("the value is %v and the address is %v\n", a, &a)
	increaseAdd(&a)
	fmt.Printf("the value is %v and the address is %v\n", a, &a)
}
func increaseAdd(add *int) {
	fmt.Printf("the value is %v before operation\n", add)
	*add++
	fmt.Printf("the value is %v before operation\n", add)
}

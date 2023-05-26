package main

import "fmt"

func main() {
	f1()
	f2("Razi")
	result3 := f3("Razi")
	fmt.Println(result3)
	x, y := f4("Razi", 30)
	fmt.Println(x)
	fmt.Println(y)
}

func f1() {
	println("this is f1 : basic function")
}

func f2(s string) {
	println("this is f2 and has  a parameter ", s)
}
func f3(s string) string {
	return "this is f3 that returns : " + s
}
func f4(s string, a int) (string, bool) {
	s = fmt.Sprint("this is f4 that returns:", s, " is ", a)
	b := true
	return s, b
}

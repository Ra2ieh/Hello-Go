package main

import "fmt"

func main() {
	var x int
	var y string
	var z bool
	x = 42
	y = "James Band"
	z = true
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

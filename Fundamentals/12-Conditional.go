package main

import "fmt"

func main() {
	x := 40
	if x == 42 {
		fmt.Println("x is 42")
	}
	fmt.Println("another thing is done")

	if x%2 != 0 {
		println("x is odd")
	} else {
		println("x is even")
	}

	if x == 0 {
		println("x is zero")
	} else if x == 41 {
		println("x is 41")
	} else {
		println("x is 40")
	}

	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}
}

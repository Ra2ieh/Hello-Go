package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" the outer loop: %d\n", i)
		for j := 1; j < 4; j++ {
			fmt.Printf("\t \t  the inner loop: %d\n", j)
		}
	}
}

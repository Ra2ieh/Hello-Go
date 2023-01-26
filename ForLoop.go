package main

import "fmt"

func main() {
	fmt.Println("print even number started...")
	loop()
	fmt.Println("print even number ended ...")
}

func loop() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			fmt.Println("******")
		}

	}
}

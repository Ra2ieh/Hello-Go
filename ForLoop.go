package main

import "fmt"

// declaration
var d = "LILI" //the scope of y would be package level
var c int      //declare a variable of type int =0
func main() {
	//short declaration
	num := 10 //declare and assign
	println(num)
	num = 99 //assign
	println(num)

	println(y)
	n, _ := fmt.Println("print even number started...")
	fmt.Println(n)
	loop()
	m, e := fmt.Println("print even number ended ...")
	fmt.Println(m)
	fmt.Println(e)
}

func loop() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			fmt.Println("******", y, "******")
		}

	}
}

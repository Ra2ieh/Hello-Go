package main

import "fmt"

// declaration
var dd = "LILI" //the scope of y would be package level
var cc int      //declare a variable of type int =0
func main() {
	println(dd)
	println(cc)
	//short declaration
	num := 10 //declare and assign
	println(num)
	num = 99 //assign
	println(num)

	n, _ := fmt.Println("print even number started...")
	fmt.Println(n)
	loop2()
	m, e := fmt.Println("print even number ended ...")
	fmt.Println(m)
	fmt.Println(e)
	loop1()
	loop3()
	loop4()
	loop5()
}

// init;condition;post
func loop2() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			fmt.Println("******")
		}

	}
}
func loop1() {
	i := 0
	for i <= 3 {
		fmt.Println("loop one :", i)
		i++
	}
}
func loop3() {
	for {
		fmt.Println("loooooooooooooooooooop")
		break
	}
}

func loop4() {
	for n := 0; n <= 7; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func loop5() {
	i := 0
	for {

		if i == 9 {
			break
		}
		fmt.Println("loop one :", i)
		i++
	}
}

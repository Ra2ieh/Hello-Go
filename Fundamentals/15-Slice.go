package main

import (
	"fmt"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Println("len is :", len(s))
	fmt.Println("cap is :", cap(s))
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println(s[3])
	fmt.Println(s[4])
	//slicing a slice
	fmt.Println("slicing a slice")
	fmt.Println(s[:])
	fmt.Println(s[0:])
	fmt.Println(s[1:])
	fmt.Println(s[2:])
	fmt.Println(s[3:])
	fmt.Println(s[4:])

	fmt.Println(s[4:])
	fmt.Println(s[3:])
	fmt.Println(s[2:4])
	fmt.Println(s[1:4])
	fmt.Println(s[0:4])
	fmt.Println(s[:])
	fmt.Println("*****************************")
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	fmt.Println("*****************************")
	for i, v := range s {
		fmt.Println(i, v)
	} //value semantic for range

	for i := range s {
		fmt.Println(i, s[i])
	} //pointer semantic for range
	//append to a slice
	fmt.Println("************** append to a slice ***************")
	s = append(s, 12, 14, 16, 18)
	fmt.Println(s)
	y := []int{100, 200, 300}
	//s = append(s, y) has error
	s = append(s, y...)
	fmt.Println(s)
	fmt.Println("************** deleting some elements ***************")
	z := append(s[:4], s[9:]...)
	fmt.Println(z)

	a := make([]int, 5, 7)
	println(a)
	a = append(a, 1)
	fmt.Println("len is :", len(a))
	fmt.Println("cap is :", cap(a))
	a = append(a, 2)
	fmt.Println("len is :", len(a))
	fmt.Println("cap is :", cap(a))
	a = append(a, 3)
	fmt.Println("len is :", len(a))
	fmt.Println("cap is :", cap(a))
	fmt.Println("************** multi-dimensional ***************")
	md := [][]int{s, y}
	fmt.Println(md)
}

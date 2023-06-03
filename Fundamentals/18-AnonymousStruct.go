package main

import "fmt"

func main() {
	//literal type
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Razi",
		last:  "Naeemi",
		age:   30,
	}
	fmt.Println(p1)
	//zero value
	p2 := struct {
		first string
		last  string
		age   int
	}{}

	fmt.Println(p2)
}

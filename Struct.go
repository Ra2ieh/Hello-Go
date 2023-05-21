package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "razi",
		last:  "naeemi",
		age:   30,
	}
	fmt.Println(p1)
	var p2 person
	p2.first = "marzi"
	p2.last = "naeemi"
	p2.age = 39
	fmt.Println(p2)

	//embedded struct

	type secretAgent struct {
		person
		ltk   bool
		first string
	}

	var sa secretAgent
	sa.ltk = true
	sa.first = "sa first"
	sa.person = person{
		first: "razi",
		last:  "naeemi",
		age:   30,
	}

	fmt.Println(sa)
	fmt.Println(sa.first, sa.person.first, sa.last, sa.age, sa.ltk)
}

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type erik struct {
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
	p3 := person{}
	fmt.Printf("%v\n", p3)
	fmt.Printf("%+v\n", p3)
	fmt.Printf("%#v\n", p3)
	last := "naeemi"
	p4 := person{
		last: last,
	}
	fmt.Printf("%#v\n", p4)
	er := erik{}
	//er=p4 has error can not convert person to erik
	p6 := struct {
		first string
		last  string
		age   int
	}{
		first: "razi",
		last:  "naeemi",
		age:   30,
	}
	er = p6 //just can covert literal struct (unnamed struct)
	fmt.Printf("%#v\n", er)
	//conversion
	er = erik(p2)
	fmt.Printf("%+v\n", er)
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

package main

import "fmt"

func main() {
	m := map[string]int{
		"Razi":  30,
		"Marzi": 40,
		"Sisi":  15,
		"pishi": 2,
	} //map[]
	fmt.Println(m)
	fmt.Println(m["Razi"])
	fmt.Println(m["Mami"])
	v, Ok := m["Mami"]
	fmt.Println(v)
	fmt.Println(Ok)

	if v, ok := m["Razi"]; ok {
		println(v)
	}
	//add an item to map
	m["dadi"] = 70
	m["mami"] = 60
	fmt.Println(m)

	//range on map

	for k, v := range m {
		fmt.Println(k, v)
	}
	//range on slice
	xi := []int{10, 12, 13, 15, 19, 21}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	//delete an item
	delete(m, "Sisi")
	delete(m, "shit") //no error

	fmt.Println(m)
	if v, ok := m["pishi"]; ok {
		fmt.Println(v)
		delete(m, "pishi")
	}
	fmt.Println(m)

}

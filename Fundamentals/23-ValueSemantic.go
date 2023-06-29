package main

import "fmt"

func main() {
	a := [4]string{"apple", "green", "orange", "yellow"}
	//pointer semantic
	for i := range a {
		if i == 1 {
			a[1] = "me"
		}
		fmt.Println(i, &a[i], a[i])

	}
	//value semantic
	fmt.Printf("\n\n")
	for i, v := range a {
		if i == 1 {
			a[1] = "you"
		}
		fmt.Println(i, &v, v)

	}
}

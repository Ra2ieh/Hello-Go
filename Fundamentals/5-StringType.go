package main

import "fmt"

func main() {
	s := "Razieh"
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("at index position of %d we have hex  %#x\n", i, v)
	}
	var a string = `razi said :

	"it is not
		fair"`

	println(a)
}

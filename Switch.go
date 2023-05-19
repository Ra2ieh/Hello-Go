package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
		/*	case (2 == 4):
				fmt.Println("this should not print2  ")

			case (4 == 4):
				fmt.Println("this should not print2  ")
				fallthrough*/
	case true:
		fmt.Println("this should  print ")
	default:
		fmt.Println("this should  print default")
	}

	switch "Bond" {
	case "Razi":
		fmt.Println("this should not print")
		/*	case (2 == 4):
				fmt.Println("this should not print2  ")

			case (4 == 4):
				fmt.Println("this should not print2  ")
				fallthrough*/
	case "Bond":
		fmt.Println("this should  print Bond")
	default:
		fmt.Println("this should  print default")
	}
	n := "Bond"
	switch n {
	case "Razi":
		fmt.Println("this should not print")
		/*	case (2 == 4):
				fmt.Println("this should not print2  ")

			case (4 == 4):
				fmt.Println("this should not print2  ")
				fallthrough*/
	case "Bond":
		fmt.Println("this should  print Bond")
	default:
		fmt.Println("this should  print default")
	}

	switch n {
	case "Razi":
		fmt.Println("this should not print")
		/*	case (2 == 4):
				fmt.Println("this should not print2  ")

			case (4 == 4):
				fmt.Println("this should not print2  ")
				fallthrough*/
	case "Bond", "Man", "To":
		fmt.Println("this should  print Bond,Man,To")
	default:
		fmt.Println("this should  print default")
	}
}

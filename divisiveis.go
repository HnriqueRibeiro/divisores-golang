package main

import "fmt"

func main() {

	for x := 1; x <= 100; x++ {

		if x%3 == 0 {

			fmt.Println(x, "é divisível por 3")

		} else {
			fmt.Println(x, "não é divisível por 3")
		}
	}
}

package main

import "fmt"

// In golang for while loop we modify for loop like while loop because in go language there is no while loop is present only 1 loop is present which is for loop
func main() {

	no := 0
	fmt.Println("Enter the number")
	fmt.Scanln(&no)
	for i := 0; i < no; i++ { // for loop in go

		fmt.Println("Pratik")
	}

	j := 0
	for j < no { // while loop in golang

		fmt.Println("Pratik")
		j++
	}

}

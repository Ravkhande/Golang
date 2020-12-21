package main

import "fmt"

func main() {

	var name [5]int = [5]int{10, 20, 30, 40, 50} // 1st way to create array
	a := [3]int{10, 20, 30}                      // 2nd way to create array

	// 3rd way to create array in golang
	var b [4]int

	b[0] = 10
	b[1] = 20
	b[2] = 30
	b[3] = 40

	// 4th way
	c := [3]int{12}
	// here length of array c is 3 but it contains only 1 element but the remaining elements of array is 0 automatically

	fmt.Println(len(name)) // len function gives the length of array

	d := []int{10, 20, 30}

	// Here to give size of array is optional it is determine by compiler to execution of program
	for i := 0; i < len(name); i++ {

		fmt.Println(name[i])
	}

}

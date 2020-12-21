package main

import "fmt"

func main() {

	no := 11
	p := &no // p is pointer which points to no

	fmt.Println("value of no : ", no)
	fmt.Println("Address of p is : ", p)
	fmt.Println("value of no through p is : ", *p)
	fmt.Println("Address of no : ", &no)

}

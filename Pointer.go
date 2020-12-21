package main

import "fmt"

func main() {

	no := 11
	p := &no

	fmt.Println(no)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&no)

}

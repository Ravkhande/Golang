package main

import "fmt"

func main() {

	choice := 0
	no1 := 11
	no2 := 23
	fmt.Println("Enter the choice")
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		Res := 0
		Res = no1 + no2
		fmt.Println("Addition is : ", Res)
		break

	case 2:
		Res := 0
		Res = no1 - no2
		fmt.Println("Subtraction is : ", Res)
		break

	case 3:
		Res := 0
		Res = no1 * no2
		fmt.Println("Multiplication is : ", Res)
		break

	case 4:
		Res := 0
		Res = no1 / no2
		fmt.Println("Division is : ", Res)
		break
	}

	marks := 0
	fmt.Println("Enter the marks of student : ")
	fmt.Scanln(&marks)

	if marks >= 60 && marks <= 70 {

		fmt.Println("First class")
	} else if marks >= 70 {

		fmt.Println("Destinction")
	} else if marks < 60 && marks >= 35 {
		fmt.Println("pass")
	}
}

/* Write a program which accept one number from user and print that number of
even numbers on screen.
Input : 7
Output: 2 4 6 8 10 12 14

*/

package main

import "fmt"

func PrintEven(no int) {

	if no <= 0 {

		no = -no
	}

	for i := 2; i <= no*2; i++ {

		if i%2 == 0 {

			fmt.Print(i, "\t")
			i++
		}

	}

}
func main() {

	no := 0
	fmt.Println("Enter the number : ")
	fmt.Scanln(&no)

	PrintEven(no)

}

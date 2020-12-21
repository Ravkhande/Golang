/* Write a program which accept number from user and print even factors of that
number.
Input : 24
Output: 1 2 4 6 8 12

*/

package main

import "fmt"

func DisplayEvenFactor(no int) {

	if no <= 0 {

		no = -no
	}

	for i := 1; i <= no/2; i++ {

		if no%i == 0 {

			if i%2 == 0 {

				fmt.Print(i, "\t")
			}

		}

	}

}
func main() {

	no := 0
	fmt.Println("Enter the number : ")
	fmt.Scanln(&no)

	DisplayEvenFactor(no)

}

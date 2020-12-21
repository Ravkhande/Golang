/*
 Accept one character from user and convert case of that character.
Input : a Output : A
Input : D Output : d
*/

package main

import "fmt"

func DisplayConvert(ch byte) {

	if ch >= 'A' && ch <= 'Z' {

		ch = ch + ('a' - 'A')
		fmt.Println(ch)

	} else if ch >= 'a' && ch <= 'z' {

		ch = ch - ('a' - 'A')
		fmt.Println(ch)
	}

}
func main() {

	var ch byte

	fmt.Println("Enter the character : ")
	fmt.Scanln(&ch)

	DisplayConvert(ch)

}

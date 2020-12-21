//  Accept one number and check whether is is divisible by 5 or not.

package main

import "fmt"

func ChkDivide(value int) bool {

	if value <= 0 {

		value = -value
	}

	if (value % 5) == 0 {

		return true
	} else {

		return false
	}
}

func main() {

	no := 0
	var bRet bool = false
	fmt.Println("Enter the number : ")
	fmt.Scanln(&no)

	bRet = ChkDivide(no)
	if bRet == true {

		fmt.Println("Number is divisible by 5")
	} else {

		fmt.Println("Number is not divisible by 5")
	}

}


//Program to divide two numbers 


package main
import "fmt"

func Divide(value1 int, value2 int) int {

	Div := 0

	if value1 <= 0 {
		value1 = -value1
	}

	if value2 <= 0 {
		value2 = -value2
	}

	Div = value1 / value2

	return Div
}
func main() {

	no1 := 0
	no2 := 0
	Result := 0
	fmt.Println("Enter the first no : ")
	fmt.Scanln(&no1)

	fmt.Println("Enter the Second no : ")
	fmt.Scanln(&no2)

	Result = Divide(no1, no2)

	fmt.Println("Division of two numbers is : ", Result)

}

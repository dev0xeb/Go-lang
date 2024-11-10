package main
import "fmt"

func main (){
	fmt.Println("Enter first number: ")
	var num1 int
	fmt.Scan(&num1)

	fmt.Println("Enter second number: ")
	var num2 int
	fmt.Scan(&num2)

	if num1 == num2{
		fmt.Println("0")
	} else if num1 > num2{
		fmt.Println("1")
	} else{
		fmt.Println("-1")
	}
}
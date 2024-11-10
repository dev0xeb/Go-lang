package main
import "fmt"

func main(){
	fmt.Println("Enetr a 5 digit number: ")
	var userInput int
	fmt.Scan(&userInput)

	if userInput < 10000 || userInput > 99999{
		fmt.Println("Enter a valid 5 digit number: ")
		return
	}

	input := userInput
	var reversedNum int
	for userInput > 0{
		remainder := userInput % 10
		reversedNum = reversedNum*10 + remainder
		userInput = userInput /10
	}

	if input == reversedNum{
		fmt.Println("number is palindrome")
	} else{
		fmt.Println("number is a palindrome")
	}
}
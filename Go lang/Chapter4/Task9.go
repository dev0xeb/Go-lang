package main
import "fmt"

func main(){
	fmt.Println("Enter a target number: ")
	var userInput int
	fmt.Scan(&userInput)

	var target int
	for target <= userInput{
		fmt.Println("Enter a number: ")
		var num int
		fmt.Scan(&num)

		target += num

		if target >= userInput{
			break
		}
	}
}
package main
import "fmt"

func main()  {
	for{
		fmt.Println("Enter a number: ")
		var userInput int
		fmt.Scan(&userInput)

		if userInput == 1 || userInput == 2{
			break
		}
	}
}
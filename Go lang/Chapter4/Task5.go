package main
import "fmt"

func main(){
	var largest = 1
	for counter :=1; counter <= 10; counter ++{
		fmt.Println("Enter a number: ")
		var number int
		fmt.Scan(&number)

		if number > largest{
			largest = number
		}
	}
	fmt.Println("The largest number enterted is: ", largest)
}
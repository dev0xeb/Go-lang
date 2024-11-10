package main
import "fmt"

func main(){
	var largest, secondLargest int
	for counter :=1; counter <= 10; counter ++{
		fmt.Println("Enter a number: ")
		var number int
		fmt.Scan(&number)

		if number > largest{
			secondLargest = largest
			largest = number
		} else if number > secondLargest{
			secondLargest = number
		}
	}
	fmt.Println("The largest number enterted is: ", largest)
	fmt.Println("The second largest number enterted is: ", secondLargest)
}
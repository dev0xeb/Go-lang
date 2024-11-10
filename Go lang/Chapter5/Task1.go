package main
import "fmt"

func main(){
	fmt.Println("How many times do you want to enter numbes: ")
	var entries int
	fmt.Scan(&entries)

	var largest, smallest int
	fmt.Println("Enter a number: ")
		var num int
		fmt.Scan(&num)

	smallest = num
	largest = num
	
	for numOfTimes := 1; numOfTimes <= entries; numOfTimes++{
		fmt.Println("Enter a number: ")
		fmt.Scan(&num)

		
		if num < smallest{
			smallest = num
		}
		if num > largest{
			largest = num
		}
		sum := largest + smallest
		fmt.Println("The smallest numeber is: ", smallest)
		fmt.Println("The largest number is: ", largest)
		fmt.Println("The sum of the largest and smallest number is: ", sum)
	}
}
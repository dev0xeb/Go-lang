package main
import "fmt"

func main() {
	var sum int
	for number := 1; number <= 30; number++{
		if number % 3 == 0{
			sum += number
		}
	}
	fmt.Println("The sum is : ", sum)
}
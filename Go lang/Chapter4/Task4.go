package main
import "fmt"

func main(){
	const taxRate1 = 0.15
	const taxRate2 = 0.20
	const sampleEarning = 30000.00

	for citizen := 1; citizen <= 3; citizen++{
		var name string
		var earnings float64
																											

		fmt.Println("Enter the name of the citizen: ", citizen)
		fmt.Scan(&name)

		fmt.Println("Enter citizen earnings: ")
		fmt.Scan(&earnings)
		
		var tax float64
		if earnings <= sampleEarning{
			tax = earnings * taxRate1
		}else {
			tax = (sampleEarning * taxRate1) + ((earnings - sampleEarning) * taxRate2)
		}
		
		fmt.Printf("The tax for %s is: %.2f\n", name, tax)
	} 
}
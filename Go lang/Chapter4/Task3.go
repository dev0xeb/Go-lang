package main
import "fmt"

func main() {
	for{
baseSalary := 200
var totalItemSold int
var totalSales float32



fmt.Println("Enter the item sold or -1 to stop: ")
var itemSold int
fmt.Scan(&itemSold)


	if itemSold == -1{
		break
	}

var itemPrice float32
switch itemSold {
case 1:
	itemPrice = 239.99
case 2:
	itemPrice = 129.75
case 3:
	itemPrice = 99.95
case 4:
	itemPrice = 350.89

default:
	
}
totalItemSold += itemSold
totalSales += float32(itemSold) * itemPrice

commission := (totalSales * 9)/100
fmt.Println("The commission for this person is: ", commission)

totalSalary := float32(baseSalary) + commission
fmt.Println("The total salary of this person is ", totalSalary)
}
fmt.Println("Exiting.......")
}
package main
import "fmt"

func main(){
	fmt.Println("Enter customer's account number: ")
	var accountNumber int
	fmt.Scan(&accountNumber)

	fmt.Println("Enter Custormer's balance at the beginning of the month: ")
	var beginningBalance int
	fmt.Scan(&beginningBalance)

	fmt.Println("Total of all items charged by customer this month: ")
	var charges int
	fmt.Scan(&charges)

	fmt.Println("Total of all credits applied by customer this month: ")
	var credits int
	fmt.Scan(&credits)

	fmt.Println("Enter allowed credit limit on customer account: ")
	var creditLimit int
	fmt.Scan(&creditLimit)

	newBalance := beginningBalance + charges - credits
	fmt.Println("Customer's new account balance is :", newBalance)

	if newBalance > creditLimit{
		fmt.Println("Credit limit exceeded.")
	}
}
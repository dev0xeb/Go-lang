package main
import "fmt"

func main(){
var totalMiles, totalGallons int
var tripCount int

for{
	fmt.Println("Enter the miles driven for this trip or enter -1 to quit: ")
	var miles int
	fmt.Scan(&miles)

	if miles == -1{
		break
	}

	fmt.Println("Enter the gallons of fuel used: ")
	var gallons int
	fmt.Scan(&gallons)

	if miles != 0 || gallons != 0{
		milesPerGallon := float64(miles) / float64(gallons)
		fmt.Printf("Miles per Gallon for this trip is: %.2f\n", milesPerGallon)
	} else { 
		fmt.Println("Miles or gallons cannot be 0")
		continue
	}

	totalMiles += miles
	totalGallons +=gallons
	tripCount++

	totalMilesPerGallon := float64(totalMiles) / float64(totalGallons)
	fmt.Printf("Total miles per gallon for all the trips is : %.2f\n", totalMilesPerGallon)
}
fmt.Println("Exiting......")
}
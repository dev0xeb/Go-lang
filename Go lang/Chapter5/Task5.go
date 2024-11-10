package main
import "fmt"

func main() {
	var gradeA, gradeB, gradeC, gradeD int

	for student := 1; student <= 5; student++ {
		fmt.Println("Enter name for student :", student)
		var name string
		fmt.Scan(&name)

		fmt.Println("Enter grade for student : ", student)
		var grade string
		fmt.Scan(&grade)

		switch grade{
		case "A":
			gradeA++
		case "B":
			gradeB++
		case "C":
			gradeC++
		case "D":
			gradeD++
			
		default:
			fmt.Println("Only grade A, B, C, D are valid")
		}
	}
	fmt.Println("The number of students with grade A:", gradeA)
	fmt.Println("The number of students with grade B:", gradeB)
	fmt.Println("The number of students with grade C:", gradeC)
	fmt.Println("The number of students with grade D:", gradeD)

}
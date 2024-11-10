package main
import "fmt"

func main(){
	var sum int64 

	fmt.Println("n\tSum")

	for n:= 1; n <= 100; n++{
		sum +=int64(n)
		fmt.Printf("%d\t%d\n", n, sum)
	}
}
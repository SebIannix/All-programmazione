package main
import "fmt"
func main() {
	var min, max int
	fmt.Println("Enter int values for the shorter and the longer sides:")
	fmt.Scan(&min, &max)
	if max<min {
		fmt.Println("Wrong values! Swapping sides...")
		min, max = max, min
	}else {
		fmt.Println("Calculating area...")
	}
	fmt.Println("Area =", max*min)
}

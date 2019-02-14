package main
import "fmt"
func main() {
	var km, mile float64
	fmt.Print("Enter kilometers: ")
	fmt.Scan(&km)
	mile = km*1.61
	var out int = int(mile)
	fmt.Println("Miles:", out)
}

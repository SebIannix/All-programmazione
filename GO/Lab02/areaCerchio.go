package main
import (
	"fmt"
)
func main() {
	var r, area float64
	fmt.Print("Enter radius:")
	fmt.Scan(&r)
	area = 3.14 * (r*r)
	fmt.Println("Area =", area)
}

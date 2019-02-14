package main
import "fmt"
func main() {
	var area, base, h int
	fmt.Print("Enter area and base of the rectangle: ")
	fmt.Scan(&area, &base)
	h = area/base
	fmt.Println("Height:", h)
}

package main
import "fmt"
func main() {
	var mum, dad, sum, media, n int
	fmt.Print("Enter your mom's age first, then your dad's: ")
	fmt.Scan(&mum, &dad)
	sum = mum + dad
	media = (mum + dad) / 2
	fmt.Println("Sum of their age:", sum)
	fmt.Println("Their average age:", media)
	fmt.Print("Wish to know the sum and average of their age years ago? Enter how many years ago in int: ")
	fmt.Scan(&n)
	sum = (mum - n) + (dad - n)
	media = ((mum - n) + (dad - n)) / 2
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", media)
}

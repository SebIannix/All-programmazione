package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	
	somma := 0
	
	for i := 1; i < 10; i++ {
		if n == 0 {
			fmt.Scan(&n)
			i--
		} else {
			somma += n
			fmt.Scan(&n)
		}
	}
	somma += n
	fmt.Println(somma)
}

package main

import (
	"fmt"
)

func fib(n int) string {
	a := ""
    b := "*"
    for i := 0; i < n; i++ {
        temp := a
        a = b
        b = temp + a
    }
    return a
}

func main() {
	var s int
	fmt.Print("Inserire numero positivo: ")
	fmt.Scan(&s)
	
	for s <= 0 {
		fmt.Print("Inserire numero positivo: ")
		fmt.Scan(&s)
	}
	
	for n := 1; n <= s; n++ {
		v := fib(n)
		fmt.Printf("%v\n", v)
	}
}

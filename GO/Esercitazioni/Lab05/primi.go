//Codice altrui ma fuck this algorithm
package main

import "fmt"

func Prime(n int, p []bool) {
	fmt.Printf("Primi inferiori a %d:\n", n)
	for i := 2; i <= n; i++ {
		if p[i] == false {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func Sieve(n int) {
	p := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = false
	}
	for i := 2; i*i <= n; i++ {
		if p[i] == false {
			for j := i*2; j <= n; j+= i {
				p[j] = true
			}
		}
	}
	Prime(n, p)
}

func main() {
	var n int
	fmt.Print("Inserire numero positivo: ")
	fmt.Scan(&n)
	
	for n <= 0 {
		fmt.Print("Inserire numero positivo: ")
		fmt.Scan(&n)
	}
	
	Sieve(n)
}

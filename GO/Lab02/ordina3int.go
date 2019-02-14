package main

import "fmt"

func main() {
	var n, m, o int
	fmt.Print("Enter 3 ints: \n")
	fmt.Scan(&n, &m, &o)
	if (n > m) && (n > o) {
		if m < o {
			m, o = o, m
		}
	} else if (m > n) && (m > o) {
		if n > o {
			n, m = m, n
		} else {
			n, m, o = m, o, n
		}
	} else if (o > n) && (o > m) {
		if n > m {
			n, m, o = o, n, m
		} else {
			n, o = o, n
		}
	}

	fmt.Println("Sequence: ", n, " >= ", m, " >= ", o)
}

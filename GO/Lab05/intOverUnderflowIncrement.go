package main

import (
	"fmt"
)

var x, n, iter int
var n32 int32
var n64 int64

func main() {
fmt.Println()
	fmt.Print("Inserire valore: ")
	fmt.Scan(&x)
	fmt.Println("Calcolo in corso...")
	
	n = x
	n32 = int32(x)
	n64 = int64(x)
	
	for {
		if (n + 1) < n {
			iter++
			n++
			n32++
			n64++
			fmt.Print("Iterazione n°", iter, ":\n")
			fmt.Println("Int:	", n, "	OVERFLOW!")
			fmt.Println("Int32:	", n32)
			fmt.Println("Int64:	", n64, "\n")
			break
		} else if (n32 + 1) < n32 {
			iter++
			n++
			n32++
			n64++
			fmt.Print("Iterazione n°", iter, ":\n")
			fmt.Println("Int:	", n)
			fmt.Println("Int32:	", n32, "	OVERFLOW!")
			fmt.Println("Int64:	", n64, "\n")
			break
		} else if (n64 + 1) < n64 {
			iter++
			n++
			n32++
			n64++
			fmt.Print("Iterazione n°", iter, ":\n")
			fmt.Println("Int:	", n)
			fmt.Println("Int32:	", n32)
			fmt.Println("Int64:	", n64, "	OVERFLOW!\n")
			break
		} else {
			iter++
			n++
			n32++
			n64++
		}
	}
}

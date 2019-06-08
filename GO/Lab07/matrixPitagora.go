package main

import(
	"fmt"
)

var m int

func main() {
	fmt.Print("Inserire dimensione della tavola: ")
	fmt.Scan(&m)
	
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			
			fmt.Print("i: ", i, "	j: ", j, "\n")
			//fmt.Print(matrix[i][j], " ")
		}
		//fmt.Println()
	}
}

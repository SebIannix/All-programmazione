package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var x int

func numScanner(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func main() {
	fmt.Print("Inserire numero: ")
	fmt.Scan(&x)
	fmt.Print("Inserire ", x, " interi:\n")
	
	var s1 []int = make([]int, x)
	
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < x && scanner.Scan(); i++ {
		y := numScanner(scanner.Text())
		s1[i] = y
	}
	fmt.Println(s1)
}

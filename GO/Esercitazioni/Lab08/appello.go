package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var args []string = os.Args[1:]
	sort.Strings(args)
	fmt.Println(args)
}

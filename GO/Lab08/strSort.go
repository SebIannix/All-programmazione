package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

type sortRunes []rune
func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Len() int {
    return len(s)
}
func sortStr(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	var in string
	fmt.Println("Inserire testo:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		in = scanner.Text()
		in += "\n"
	}
	
	out := sortStr(in)
	fmt.Println(out)
}

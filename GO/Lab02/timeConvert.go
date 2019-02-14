package main
import "fmt"
func main() {
	var in, h, m, s float64
	fmt.Println("Select conversion: Hour to Min. (1), Min. to Hour (2), Min. to Sec. (3), Sec. to Min. (4)")
	fmt.Scan(&in)
	if in == 1 {
		fmt.Print("Enter hours to convert to minutes: ")
		fmt.Scan(&h)
		m = h * 60
		fmt.Println(h, "hours are", m, "minutes.")
	} else if in == 2 {
		fmt.Print("Enter minutes to convert to hours: ")
		fmt.Scan(&m)
		h = m / 60
		fmt.Println(m, "minutes are", h, "hours.")
	} else if in == 3 {
		fmt.Print("Enter minutes to convert to seconds: ")
		fmt.Scan(&m)
		s = m * 60
		fmt.Println(m, "minutes are", s, "seconds.")
	} else if in == 4 {
		fmt.Print("Enter seconds to convert to minutes: ")
		fmt.Scan(&s)
		m = s / 60
		fmt.Println(s, "seconds are", m, "minutes.")
	} else {
		fmt.Println(in, "is not bound. Exiting program...")
	}
}

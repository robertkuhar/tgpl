package main
import "fmt"

func main() {
	months := [...]string{
		1: "Jan", 2: "Feb", 3: "Mar",
		4: "Apr", 5: "May", 6: "Jun",
		7: "Jul", 8: "Aug", 9: "Sep",
		10: "Oct", 11: "Nov", 12: "Dec"}
	fmt.Printf("months %T %v\n", months, months)
	// Wait, what?  [13]string [ Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec]

	bills := [...]string{1: "Geo", 5: "Abe", 10: "Alex", 20: "Tom"}
	fmt.Printf("bills %T %v\n", bills, bills)
	// Wow:  [21]string [ Geo    Abe     Alex          Tom]

	Q1 := months[1:4]
	Q2 := months[4:7]
	Q3 := months[7:10]
	Q4 := months[10:]
	fmt.Printf("Q1 %T %v\n", Q1, Q1)
	fmt.Printf("Q2 %T %v\n", Q2, Q2)
	fmt.Printf("Q3 %T %v\n", Q3, Q3)
	fmt.Printf("Q4 %T %v\n", Q4, Q4)
}

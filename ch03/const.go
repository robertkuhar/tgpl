package main
import "fmt"

type Bob int

const (
	First Bob = 1
	Second = First + 1
	Third = Second + 1
)

func main() {
	fmt.Printf("First %d, Second %d, Third %d", First, Second, Third)
}

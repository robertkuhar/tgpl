package main

import (
	"time"
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func EmployeeByID(id int) *Employee {
	return &dilbert
}

func otherEmployeeById(id int) Employee {
	return dilbert
}

type Point struct{ X, Y int }

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

func main() {
	var 世界 = 1
	fmt.Printf("hey %v\n", 世界)
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	dilbert.Position = "Assistand Under Secretary"

	position := &dilbert.Position
	dilbert.Name = "Dilbert"
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	fmt.Printf("dilbert: %v\n\n", dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Printf("employeeOfTheMonth: %v\n\n", *employeeOfTheMonth)

	var empPtr *Employee = EmployeeByID(dilbert.ID)
	empPtr.Salary = 10
	fmt.Printf("EmployeeById(%d): %v\n", dilbert.ID, *empPtr)
	fmt.Printf("dilbert: %v\n\n", dilbert)

	// By what sorcery does dilbert end up unchanged here?  Some sort of copy constructor?
	var emp = otherEmployeeById(dilbert.ID)
	emp.Salary = 20
	fmt.Printf("otherEmployeeById(%d): %v\n", dilbert.ID, emp)
	fmt.Printf("dilbert: %v\n\n", dilbert)
	fmt.Printf("&dilbert: %p, *empPtr: %p, &emp: %p\n\n", &dilbert, empPtr, &emp)

	for i := 0; i < 3; i++ {
		var tmp = otherEmployeeById(dilbert.ID)
		tmp.Salary *= (i + 1)
		fmt.Printf("&dilbert: %p Salary: %d, &tmp: %p Salary: %d\n", &dilbert, dilbert.Salary, &tmp, tmp.Salary)
	}

	p1_2 := Point{1, 2}
	scaledPoint := Scale(p1_2, 5)
	fmt.Printf("p1_2: %v, scaledPoint: %v\n\n", p1_2, scaledPoint)

	p1_2_a := Point{1, 2}
	p1_2_b := Point{1, 2}
	p2_2_a := Point{2, 2}
	fmt.Printf("%v==%v %v\n", p1_2_a, p1_2_b, p1_2_a == p1_2_b)
	fmt.Printf("%v==%v %v\n", p1_2_a, p2_2_a, p1_2_a == p2_2_a)
}

package main

import (
	"fmt"
	"math"
)

// Shape 接口
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	length, width float64
}

func (rect *Rectangle) Area() {
	area := rect.length * rect.width
	fmt.Println("Rectangle Area:", area)
}

func (rect *Rectangle) Perimeter() {
	perimeter := (rect.length + rect.width) * 2
	fmt.Println("Rectangle Perimeter:", perimeter)
}

type Circle struct {
	radius float64
}

func (circle *Circle) Area() {
	area := math.Pi * circle.radius * circle.radius
	fmt.Println("Circle Area:", area)
}

func (circle *Circle) Perimeter() {
	perimeter := math.Pi * circle.radius * 2
	fmt.Println("Circle Perimeter:", perimeter)
}

// Person  结构体
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (emp Employee) PrintInfo() {
	fmt.Println("Employee ID:", emp.EmployeeID)
	fmt.Println("Employee Name:", emp.Name)
	fmt.Println("Employee Age:", emp.Age)
}

func main() {

	fmt.Println("*****************面向对象***************")

	//OOP1
	rectangle := Rectangle{length: 3, width: 4}
	var shape1 Shape = &rectangle
	shape1.Area()
	shape1.Perimeter()

	circle := Circle{radius: 2}
	var shape2 Shape = &circle
	shape2.Area()
	shape2.Perimeter()

	//OOP2
	emp := Employee{Person{"小王", 28}, "123"}
	emp.PrintInfo()

}

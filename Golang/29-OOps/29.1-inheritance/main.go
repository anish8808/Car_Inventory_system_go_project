package main

import "fmt"

//in go we dont have inherictance but we can achive inheritance by using compisition

type Person struct {
	Name string
	Age  int
}

func (p *Person) Intro() {
	fmt.Println(p.Age, p.Name)
}

type Employee struct {
	Person     //extending the Person strcut from Employee
	EmployeeID int
}

func (e Employee) Intro() {
	fmt.Println(e.Name, e.EmployeeID, e.Age) // here we can use name age from Person
}
func main() {
	p := Person{Name: "Anish", Age: 25}
	ee := Employee{Person: p, EmployeeID: 123}
	ee.Person.Intro()
	ee.Intro()
	p.Intro()
}

package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c.Age)
	fmt.Println(c.Name)
	fmt.Println(c.ModelNo)
}

func main() {
	c := Car{
		Name:    "Volvo",
		Age:     0,
		ModelNo: 2021,
	}
	c.Print()
}

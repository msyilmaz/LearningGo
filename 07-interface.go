package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("start")
	Anything(2.44)
	Anything("Safa")
	Anything(struct {
		Name string
		Age  int
	}{
		"Safa",
		27,
	})
}

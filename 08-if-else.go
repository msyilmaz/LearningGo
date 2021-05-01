package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Value is nill")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}

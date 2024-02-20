package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var ptr *int
	fmt.Println(ptr)
	var1 := 23
	varptr := &var1
	fmt.Println("Memory", varptr)
	fmt.Println("Value", *varptr)
	*varptr = *varptr + 2
	fmt.Println(var1 == *varptr)

}

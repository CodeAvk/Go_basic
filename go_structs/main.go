package main

import "fmt"

func main() {
	fmt.Println("Go")
	// no class in go instea it has struct
	// No inheritnce in golang,No super or parent
	avk := User{"Abhisek", "avksmlavk@gmail.com", 23, true}
	fmt.Println(avk)
	fmt.Printf("%+v\n", avk)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

package main

import "fmt"

func main() {
	fmt.Println("Go")
	user1 := User{1, "Abhisek", false}
	user1.IsActive()
	fmt.Println("before", user1.Id)
	user1.updateId(8)
	fmt.Println("after", user1.Id)
}

type User struct {
	Id     int
	Name   string
	Status bool
}

func (u User) IsActive() {
	fmt.Println("Is user active :", u.Status)
}

func (u *User) updateId(i int) {
	u.Id = i
	fmt.Println("New id is ", u.Id)
}

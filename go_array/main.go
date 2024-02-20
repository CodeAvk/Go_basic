package main

import "fmt"

func main() {
	fmt.Println("Ok")
	var fruitList [4]string
	fruitList[0] = "A"
	fruitList[1] = "A1"
	fruitList[3] = "A2"
	fmt.Println("The Fruitlist is :", fruitList)
	fmt.Println("The Fruitlist length is :", len(fruitList))
	var vegList = [3]string{"P", "T", "2"}
	fmt.Println("veglist is :", vegList)
}

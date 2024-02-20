package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	wc := "XYZ"
	fmt.Println(wc)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for Rating, ", input)

}

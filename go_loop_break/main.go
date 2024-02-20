package main

import "fmt"

func main() {
	fmt.Println("Hi")
	days := []string{"Sunday", "MON", "TUe", "WEd"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
}

package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://pro.learncodeonline.in/learn/subscription/LCO-Pro-12-months/LCO-Pro-6-month?courseId=106114"

func main() {
	fmt.Println("IK")
	fmt.Println(myurl)

	// parse
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("1", result.Scheme)
	fmt.Println("2", result.Host)
	fmt.Println("3", result.Path)
	fmt.Println("4", result.Port())
	fmt.Println("5", result.RawQuery)

}

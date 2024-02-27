package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://abhishek-portfolio-scarr33.vercel.app/"

func main() {
	fmt.Println("Go")

	// Create a custom HTTP client that skips certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get(url)

	if err != nil {
		panic(err)
	}
	defer fmt.Printf("Response type is : %T\n", response)

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	println(content)

}

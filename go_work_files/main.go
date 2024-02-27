package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("JI")
	data := "Hi Abhisek is Here 👋🏻"
	file, err := os.Create("./my-file.txt")
	checkError(err)
	defer file.Close() // Close the file when main() exits

	length, err := io.WriteString(file, data)
	checkError(err)
	fmt.Println("length is :", length)

	readFile("./my-file.txt")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func abc() {
	fmt.Println("ntg1")
}
func readFile(file string) {
	bytedata, err := ioutil.ReadFile(file)
	checkError(err)
	fmt.Println("byte-data", string(bytedata))
}

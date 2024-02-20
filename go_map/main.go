package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"
	languages["rb"] = "ruby"
	fmt.Println(languages)
	delete(languages, "rb")
}

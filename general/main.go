package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("lets make folder in temp folder at /users/ashit.vyas/temp")
	os.Mkdir("/Users/ashit.vyas/temp/goLogs", os.ModePerm)
}

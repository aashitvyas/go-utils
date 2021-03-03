package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func dircheck(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	if len(files) == 0 {
		fmt.Println("Directory is empty")
	} else {
		for _, file := range files {
			fmt.Println("File Name is", file.Name())
		}
		fmt.Println("Total number of files found", len(files))
	}
}

func main() {

	dirPath := flag.String("dir", "foo", "a directory path")
	flag.Parse()
	fmt.Printf("This program is going to check if %s is empty or not\n", *dirPath)
	dircheck(*dirPath)
	fmt.Println("Done!")

}

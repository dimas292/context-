package main

import (
	"fmt"
	"os"
)

func main() {

	names := []string{"Nobee", "bee", "go", "Nodejs"}

	for _, name := range names {
		if name == "go"{
			os.Exit(1)
		}
		cetak(name)
	}
}

func cetak(text string){
	fmt.Println(text)
}


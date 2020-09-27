package main

import (

	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, argumento :=  range os.Args[1:] {

		s+=sep + argumento
		sep=" "
	}
	fmt.Println(s)
}

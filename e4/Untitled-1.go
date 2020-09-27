package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Diretorio do programa copilado: " ,os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], "|"))
}

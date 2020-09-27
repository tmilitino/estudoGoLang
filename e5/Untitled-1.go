package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Diretorio do programa copilado: " ,os.Args[0])
	count :=make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan(){
	
		if input.Text()=="exit" {
			fmt.Println("fim da entrada")
			break
		}
		count[input.Text()] = len(count) + 1
	}

	input.Err()
	for line, n := range count{
		fmt.Printf("%d\t%s\n",n,line)
	}
}

package main

import (
	"fmt"
	"os"
	"bufio"
)

func countlines(f *os.File, count map[string]int){
	input :=bufio.NewScanner(f)
	for input.Scan(){
		count[input.Text()]++
	}
	input.Err()
}

func main(){
count := make(map[string]int)
files := os.Args[1:]
if len(files)==0{
	countlines(os.Stdin, count)
}else{
for _, argumento := range files{
	f, err := os.Open(argumento)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		continue
	}
	countlines(f, count)
	f.Close()
}
}
for line, n := range count {
	fmt.Printf("%d\t%s\n",n,line)
}
}
package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr, "connect Url: %v\n", err)
			os.Exit(1)
		}
		bodyRead, errRead := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if errRead!=nil{
			fmt.Fprintf(os.Stderr, "read body: %v\n", errRead)
			os.Exit(1)
		}
		fmt.Printf("%s",bodyRead)
	}
}
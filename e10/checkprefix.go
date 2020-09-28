package main

import(
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	for _, url := range os.Args[1:]{
		if !(strings.HasPrefix(url, "http")){
			url ="http://"+url
		}
		resp, err := http.Get(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr, "connect Url: %v\n", err)
			os.Exit(1)
		}
		bodyRead, errRead := io.Copy(os.Stdout ,resp.Body)
		statusRead := resp.Status

		resp.Body.Close()
		if errRead!=nil{
			fmt.Fprintf(os.Stderr, "read body: %v\n", errRead)
			os.Exit(1)
		}
		fmt.Println(bodyRead)
		fmt.Println(statusRead)
	}
}
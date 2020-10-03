package main

import (
	"fmt"
)

type jogadorInfo struct{
	ID string `json:"ID"`
	jogadorID string `json:"PlayerID"`
	nome string `json:"Name"`
	idade int `json:"Age"`
}
func main(){
	lista := make(chan []jogadorInfo	)
fmt.Println(lista)
}
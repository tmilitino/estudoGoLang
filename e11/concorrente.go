package main

import(
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	start := time.Now()

	ch :=make(chan string)

	for _, url := range os.Args[1:]{
		go fetch(url, ch)

	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


func fetch(url string, ch chan<- string) {
		start := time.Now()

		reps, errGet := http.Get(url)
		if errGet != nil {
			ch <- fmt.Sprint(errGet)
			return
		}
		nbytes, errCopy:= io.Copy(ioutil.Discard,reps.Body)
		reps.Body.Close()

		if errCopy != nil {
		ch <- fmt.Sprint("teste")
		return
		}

	secs := time.Since(start).Seconds()
	ch <-	fmt.Sprintf("%.2fs %7d %s", secs,nbytes, url)	
}
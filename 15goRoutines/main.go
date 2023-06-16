package main

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go fetch("https://www.golangprograms.com")
	go fetch("https://github.com/")
	go fetch("https://stackoverflow.com")
	wg.Wait()
}

func fetch(url string) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	println(body, url)

}

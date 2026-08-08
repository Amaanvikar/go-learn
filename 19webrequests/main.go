package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {

	fmt.Println("webrequests in go")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}

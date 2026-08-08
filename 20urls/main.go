package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://localhost:8080/google.com/learn?coursename=reactjs&paymentid=ghp_1234567890"

func main() {

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Query())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost:8080",
		Path:     "/google.com/learn",
		RawQuery: "coursename=reactjs&paymentid=ghp_1234567890",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}

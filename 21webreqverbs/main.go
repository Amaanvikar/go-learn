package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostformRequest()
}

func PerformGetRequest() {

	myurl := "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)

	bytecount, _ := responseString.Write(content)

	fmt.Println(string(content))

	fmt.Println(bytecount)

	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {

	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
	"Coursename": "learn with go-lang",
	"price": "0",
	"platform": "golangfornocost.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// 	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	// bytecount, _ := responseString.Write(content)

	// fmt.Println(string(content))

	// fmt.Println(bytecount)

	// fmt.Println(responseString.String())

	fmt.Println(content)
}

func PerformPostformRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Aman")
	data.Add("lastname", "pathan")
	data.Add("age", "24")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

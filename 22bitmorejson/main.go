package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json encoding in go-lang")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"Reactjs Bootcamp", 299, "reactaman.dev", "abc124", []string{"web-dev", "js"}},
		{"Golang backend", 199, "golangaman.dev", "xyz124", []string{"backend", "go"}},
		{"Mernstack", 299, "mernaman.dev", "124abc", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	
	{
	"coursename": "Reactjs Bootcamp",
	"Price": 299,
	"website": "reactaman.dev",
	"Password": "abc124",
	"Tags": ["web-dev","js"]
	}
	`)

	var courses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Valid Json")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("Invalid Json")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("%#\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}

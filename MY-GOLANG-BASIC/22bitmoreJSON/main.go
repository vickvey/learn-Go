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
	fmt.Println("Welcome to JSON video!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "google.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "google.com", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "google.com", "aly123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "google.com",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is: %v and value is: %v and Type is: %T\n", key, value, value)
	}

}

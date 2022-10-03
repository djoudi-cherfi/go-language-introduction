package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func jsonOne() {
	jsonFromApi := `
	[
		{
			"name": "John",
			"age": 25,
			"email": "john@foo.com",
			"phone": "0123456789"
		},
		{
			"name": "Jane",
			"age": 20,
			"email": "jane@bar.com",
			"phone": "0987654321"
		}
	]
	`

	var users []User

	err := json.Unmarshal([]byte(jsonFromApi), &users)

	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}

	fmt.Printf("json: %v\n", users)

	//  ----------------------------------

	var myStruct []User

	var userOne User
	userOne.Name = "Chris"
	userOne.Age = 35
	userOne.Email = "chris@baz.com"
	userOne.Phone = "0432198765"

	myStruct = append(myStruct, userOne)

	var userTwo User
	userTwo.Name = "Kate"
	userTwo.Age = 40
	userTwo.Email = "kate@quux.com"
	userTwo.Phone = "9876504321"

	myStruct = append(myStruct, userTwo)

	jsonFromSlice, err := json.MarshalIndent(myStruct, "", " ")

	if err != nil {
		fmt.Println("Error marshalling:", err)
	}

	fmt.Println(string(jsonFromSlice))
}

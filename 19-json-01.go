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
			"email": "John@foo.com",
			"phone": "0123456789"
		},
		{
			"name": "Jane",
			"age": 20,
			"email": "Jane@bar.com",
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
}

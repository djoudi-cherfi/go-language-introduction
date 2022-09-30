package main

import "fmt"

type Contact struct {
	Name  string
	Age   int
	Phone map[string]string
}

func newContact(name string, age int, phone map[string]string) Contact {
	c := Contact{
		Name:  name,
		Age:   age,
		Phone: phone,
	}

	return c
}

func (c Contact) displayContact() string {
	display := fmt.Sprintf("Contact: %v (%v ans)\n", c.Name, c.Age)

	display += "----------------\n"

	for key, value := range c.Phone {
		display += fmt.Sprintf("%v: %v\n", key, value)
	}

	return display
}

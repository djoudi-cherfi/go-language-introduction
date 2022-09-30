package main

import "fmt"

type contact struct {
	name  string
	age   int
	phone map[string]string
}

func newContact(name string, age int, phone map[string]string) contact {
	c := contact{
		name:  name,
		age:   age,
		phone: phone,
	}

	return c
}

func (c contact) displayContact() string {
	display := fmt.Sprintf("Contact: %v (%v ans)\n", c.name, c.age)

	display += "----------------\n"

	for key, value := range c.phone {
		display += fmt.Sprintf("%v: %v\n", key, value)
	}

	return display
}

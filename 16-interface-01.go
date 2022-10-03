package main

import "fmt"

type Cat struct {
	Name  string
	Breed string
}

type Spider struct {
	Name     string
	Breed    string
	Venomous bool
}

type Animal interface {
	Noise() string
	NumberOfLegs() int
}

func (c *Cat) Noise() string {
	return "miaou"
}

func (c *Cat) NumberOfLegs() int {
	return 4
}

func (s *Spider) Noise() string {
	return "hiss"
}

func (s *Spider) NumberOfLegs() int {
	return 8
}

func printAnimalInfo(a Animal) {
	fmt.Println("Le bruit de cet animal est", a.Noise(), "est il possède", a.NumberOfLegs(), "pattes !")
}

func interfaceOne() {
	catOne := Cat{
		Name:  "Kitty",
		Breed: "Siamois",
	}

	printAnimalInfo(&catOne)

	spiderOne := Spider{
		Name:     "Spi",
		Breed:    "Veuve noir",
		Venomous: true,
	}

	printAnimalInfo(&spiderOne)

}

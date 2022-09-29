package main

import (
	"fmt"
	"sort"
)

func mapOne() {
	/*
		variable := map[type_key] type_value{
			"key": "value",
		}
	*/
	supermarketPrice := map[string]int{
		"gateau": 8,
		"eau":    2,
		"viande": 6,
	}

	fmt.Println("Le prix est de :", supermarketPrice["gateau"])

	fmt.Println("----------- supermarketPrice")

	for key, value := range supermarketPrice {
		fmt.Println(key, value)
	}

	daysInYear := 0

	daysInMonth := map[string]int{
		"january":   31,
		"february":  28,
		"march":     31,
		"april":     30,
		"may":       31,
		"june":      30,
		"july":      31,
		"august":    31,
		"september": 30,
		"october":   31,
		"november":  30,
		"december":  31,
	}

	for _, value := range daysInMonth {
		daysInYear += value
	}

	fmt.Println("----------- daysInYear")
	fmt.Printf("Nombre de jours dans une année : %d jours !\n", daysInYear)

	fmt.Println("----------- daysInMonth")
	for key, value := range daysInMonth {
		fmt.Printf("%v possède %d jours !\n", key, value)
	}

	fmt.Println("----------- order daysInMonth")
	keys := make([]string, 0)
	for k := range daysInMonth {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, daysInMonth[k])
	}
}

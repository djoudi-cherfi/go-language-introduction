package main

import (
	"fmt"
	"tuto/go/function"
)

// Fonction qui retourne un nombre aléatoire selon une valeur donnée
func CalculateValue(intChan chan int) {
	randomNumber := function.GenerateRandomNumber(50)

	intChan <- randomNumber
}

func channelOne() {
	// On ouvre le channel pour envoyer et recevoir des données du type précisé
	foo := make(chan int)

	// On ferme le channel
	defer close(foo)

	// Appel de la fonction en une goroutine
	go CalculateValue(foo)

	// On enregistrer les données reçues
	num := <-foo

	// On affiche les données
	fmt.Println(num)
}

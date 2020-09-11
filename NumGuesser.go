package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var help int
	random := rand.Intn(100)
	var user int
	rules := "Le but est simple, le programme choisit un nombre aléatoire entre 0 et 100. A toi de le devinez."

	for {
		fmt.Println("Votre choix: ")
		fmt.Scan(&user)
		if user == help {
			fmt.Println(rules)
		}
		if user < random {
			fmt.Print("C'est plus!\n")
		}
		if user > random {
			fmt.Println("C'est moins!")
		}
		if user == random {
			fmt.Println("GG")
			break
		}
		if user >100 {
			fmt.Println("Number out of limits or unexpected characters find, please try again.")
		}
		if user < 0 {
			fmt.Println("Number out of limits or unexpected characters find, please try again.")
		}
	}
}
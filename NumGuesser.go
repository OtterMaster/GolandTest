package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	random := rand.Intn(100)
	var user int


	for {
		fmt.Println("Votre choix: ")
		fmt.Scan(&user)
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
	}
}
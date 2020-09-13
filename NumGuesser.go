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
		fmt.Println("Your choice: ")
		fmt.Scan(&user)
		
		if user < 0 {
			fmt.Print("Number out of limits or unexpected characters find, please try again.\n")
		} else if user < random {
			fmt.Println("Higher than your choice!")
		}
		if user > 100 {
			fmt.Println("Number out of limits or unexpected characters find, please try again.")
		} else if user > random {
			fmt.Println("Lower than your choice!")
		}

		if user == random {
			fmt.Println("GJ boi!")
			break
		}
	}
}
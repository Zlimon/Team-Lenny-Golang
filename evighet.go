package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomteller := 0
	for i := 0; i < 1; randomteller++ {
		random := rand.Intn(1000000)
		randomstopper := 50000
		fmt.Println("Et tall:", random)
		if random == randomstopper {
			fmt.Println("WoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWo")
			fmt.Println(randomteller, "forsøk ble gjort før" ,randomstopper, "ble nådd!")
			fmt.Println("WoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWoWo")
			i++
		}
	}
}
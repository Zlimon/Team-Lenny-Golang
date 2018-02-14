package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 1; {
		random := rand.Intn(1000000)
		fmt.Println("Et tall:", random)
	}
}

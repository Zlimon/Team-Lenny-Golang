package main

import (
	"log"
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	fil, err := os.Open("C:/Users/Simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave2/text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fil2, err := ioutil.ReadFile("C:/Users/Simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave2/text1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(fil)

	/**
	En løkke som benytter fileScanner for å telle hver linje.
	 */
	lineCounter := 0
	for fileScanner.Scan() {
		lineCounter++
	}

	s := string(fil2) //Konverterer text.txt til en string.

	amountRuneA := strings.Count(s, "A") + strings.Count(s, "a")
	amountRuneB := strings.Count(s, "B") + strings.Count(s, "b")
	amountRuneC := strings.Count(s, "C") + strings.Count(s, "c")
	amountRuneD := strings.Count(s, "D") + strings.Count(s, "d")
	amountRuneE := strings.Count(s, "E") + strings.Count(s, "e")
	amountRuneF := strings.Count(s, "F") + strings.Count(s, "f")
	amountRuneG := strings.Count(s, "G") + strings.Count(s, "g")
	amountRuneH := strings.Count(s, "H") + strings.Count(s, "h")
	amountRuneI := strings.Count(s, "I") + strings.Count(s, "i")
	amountRuneJ := strings.Count(s, "J") + strings.Count(s, "j")
	amountRuneK := strings.Count(s, "K") + strings.Count(s, "k")
	amountRuneL := strings.Count(s, "L") + strings.Count(s, "l")
	amountRuneM := strings.Count(s, "M") + strings.Count(s, "m")
	amountRuneN := strings.Count(s, "N") + strings.Count(s, "n")
	amountRuneO := strings.Count(s, "O") + strings.Count(s, "o")
	amountRuneP := strings.Count(s, "P") + strings.Count(s, "p")
	amountRuneQ := strings.Count(s, "Q") + strings.Count(s, "q")
	amountRuneR := strings.Count(s, "R") + strings.Count(s, "r")
	amountRuneS := strings.Count(s, "S") + strings.Count(s, "s")
	amountRuneT := strings.Count(s, "T") + strings.Count(s, "t")
	amountRuneU := strings.Count(s, "U") + strings.Count(s, "u")
	amountRuneV := strings.Count(s, "V") + strings.Count(s, "v")
	amountRuneW := strings.Count(s, "W") + strings.Count(s, "w")
	amountRuneX := strings.Count(s, "X") + strings.Count(s, "x")
	amountRuneY := strings.Count(s, "Y") + strings.Count(s, "y")
	amountRuneZ := strings.Count(s, "Z") + strings.Count(s, "z")

	amountRunes := []int {
		amountRuneA, amountRuneB, amountRuneC, amountRuneD, amountRuneE,
		amountRuneF, amountRuneG, amountRuneH, amountRuneI, amountRuneJ,
		amountRuneK, amountRuneL, amountRuneM, amountRuneN, amountRuneO,
		amountRuneP, amountRuneQ, amountRuneR, amountRuneS, amountRuneT,
		amountRuneU, amountRuneV, amountRuneW, amountRuneX, amountRuneY, amountRuneZ,
	}

	var max int = amountRunes[0]
	for _, value := range amountRunes {
		if max < value {
			max = value
		}
	}
	fmt.Println(max)

	fmt.Println("Information about text.txt:")
	fmt.Println("Number of lines in file:", lineCounter)
	//fmt.Println("Total runes:", utf8.RuneCount(fil2))
	fmt.Println("Most common runes:")
	fmt.Println("1. Rune:", amountRuneA, amountRunes[0])
	fmt.Println("2. Rune:", amountRuneB, amountRunes[1])
	fmt.Println("3. Rune:", amountRuneC)
	fmt.Println("4. Rune:", amountRuneD)
	fmt.Println("5. Rune:", amountRuneE)
}
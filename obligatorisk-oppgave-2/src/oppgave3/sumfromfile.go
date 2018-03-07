package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

var i int

func main() {
	fil, err := ioutil.ReadFile("C:/Users/simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave3/file.txt")
	errorCheck(err)

	str := string(fil) //Konverterer filen til en string.

	if strings.Contains(str, "5") {
		fmt.Println("Found a number! \n", str)
	} else {
		fmt.Println("Found nothing")
	}
}


//fmt.Sscanf(str, "nr1:", &i)
//fmt.Println(i) // Outputs 123

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
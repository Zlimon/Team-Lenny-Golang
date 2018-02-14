//Oppgave 4
package main

import(
	"fmt"

)
func main() {
	for j := 0x80; j <= 0xFF; j++ {
		fmt.Printf("%X  ", j)
		fmt.Printf("%c  ", j)
		fmt.Printf("%b  ", j)
		fmt.Println("")
	}
}

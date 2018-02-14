/**Oppgave 4
* 0x80 - 0x9F blir vist som en firekant istedenfor sitt satte symbol på alle datamaskinene undtatt en macbook
* hvor 0x80 blir vist som deletegnet. 
* Dette er nok på grunn av forskjellig standardvalgt file encoding på windows og macOS
*
*
**/
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

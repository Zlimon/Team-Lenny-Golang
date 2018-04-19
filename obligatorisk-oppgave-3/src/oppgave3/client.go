package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"regexp"
)

var (
	connHost = "127.0.0.1"
	connPort = "17"
	connType = "tcp"
)

var yesAnswer = regexp.MustCompile("y|Y|yes|Yes|YES")
var noAnswer = regexp.MustCompile("n|N|no|No|NO")

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to connect to the server", connHost,"at port", connPort, "? Y for yes, N for no...")
	answer, _ := reader.ReadString('\n')

	if yesAnswer.MatchString(answer) {
		// Connect to the server
		conn, err := net.Dial(connType, connHost+":"+connPort)
		if err != nil {
			fmt.Println("Error dialing:", err.Error())
			os.Exit(1)
		}

		fmt.Fprintf(conn, answer + "\n")
		quote, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Quote from server:\n"+quote)
		conn.Close()
	} else if noAnswer.MatchString(answer) {
		fmt.Println("Neivel!")
	} else {
		fmt.Print("You have to answer either Y or N! Not: ", answer)
	}
}
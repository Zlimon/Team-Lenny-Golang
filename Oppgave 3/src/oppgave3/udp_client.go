package oppgave3

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
	connTCP = "tcp"
	connUDP = "udp"
)

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

var yesAnswer = regexp.MustCompile("y|Y|yes|Yes|YES")
var noAnswer = regexp.MustCompile("n|N|no|No|NO")

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to connect to the server", connHost,"at port", connPort, "? Y for yes, N for no...")
	answer, _ := reader.ReadString('\n')

	if yesAnswer.MatchString(answer) {
		udpConn, err := net.Dial(connUDP, connHost+":"+connPort)
		CheckError(err)

		for {
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(udpConn, text + "\n")
			message, _ := bufio.NewReader(udpConn).ReadString('\n')
			fmt.Print("Message from server: "+message)
		}
	} else if noAnswer.MatchString(answer) {
		fmt.Println("Neivel!")
	} else {
		fmt.Print("You have to answer either Y or N! Not: ", answer)
	}
}
package oppgave3

import (
	"net"
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"
)

var (
	connHost = "127.0.0.1"
	connPort = "17"
	connTCP = "tcp"
	connUDP = "udp"
)

var quotes = []string {
	"Don’t text me while i’m in the middle of texting you. Now i have to change my whole text.",
	"When you become really close to someone, you can hear their voice in your head when you read their texts.",
	"People make time for who they want to make time for. People text, call and reply to people they want to talk to. Never believe anyone who says they’re too busy. If they wanted to be around you, they would.",
	"I just want someone who won’t get annoyed when I text them six times or in all caps. Someone I can go on long drives with and can sing along to the radio with. Someone I can eat pizza with at 2am and kiss at 6pm. Someone who chooses me everyday and never thinks twice about it."}

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

func main() {
	tcpConn, err := net.Listen(connTCP, connHost + ":" + connPort)
	CheckError(err)

	defer tcpConn.Close()

	fmt.Println("Listening on", connHost, ":", connPort, "with", connTCP)

	for {
		tcpConn, err := tcpConn.Accept()
		CheckError(err)

		go func(net.Conn) {
			rand.Seed(time.Now().UnixNano())
			random := rand.Intn(3)

			fmt.Println("Recieved a connection from")
			fmt.Println("Returning a daily quote...")
			tcpConn.Write([]byte(quotes[random] + "\n"))
			fmt.Println("Connection closed.")
		}(tcpConn)

		udpAddr, err := net.ResolveUDPAddr(connUDP, ":"+connPort)
		CheckError(err)

		udpConn, err := net.ListenUDP(connUDP, udpAddr)
		CheckError(err)

		defer udpConn.Close()

		fmt.Println("Listening on", connHost, ":", connPort, "with", connUDP)

		buf := make([]byte, 1024)

		for {
			n, addr, err := udpConn.ReadFromUDP(buf)
			CheckError(err)

			fmt.Println("Received", string(buf[0:n]), "from", addr)

			newMessage := strings.ToUpper(string(buf[0:n]))
			udpConn.WriteToUDP([]byte(newMessage+"\n"), addr)
		}
	}
}
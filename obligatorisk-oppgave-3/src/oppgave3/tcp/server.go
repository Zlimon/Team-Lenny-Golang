package main

import (
	"net"
	"fmt"
	"os"
	"strings"
	"math/rand"
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
	tcpConn, err := net.Listen(connTCP, connHost+":"+connPort)
	udpAddr,err := net.ResolveUDPAddr(connUDP,":"+connPort)
	CheckError(err)

	udpConn, err := net.ListenUDP(connUDP, udpAddr)
	CheckError(err)

	defer tcpConn.Close()
	defer udpConn.Close()
	fmt.Println("Listening on ",connHost,":",connPort)

	buf := make([]byte, 1024)

	for {
		tcpConn, err := tcpConn.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		} else {
			random := rand.Intn(3)

			fmt.Println("Recieved a connection!")
			fmt.Println("Returning a daily quote...")
			tcpConn.Write([]byte(quotes[random] + "\n"))
			fmt.Println("Connection closed.")
		}

		for {
			n,addr,err := udpConn.ReadFromUDP(buf)
			CheckError(err)
			fmt.Println("Received",string(buf[0:n]), "from",addr)

			newmessage := strings.ToUpper(string(buf[0:n]))
			udpConn.WriteToUDP([]byte(newmessage + "\n"),addr)
		}
	}
}
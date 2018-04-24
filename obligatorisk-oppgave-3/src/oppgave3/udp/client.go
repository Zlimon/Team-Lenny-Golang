package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "localhost:10001")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
	conn.Close()
}
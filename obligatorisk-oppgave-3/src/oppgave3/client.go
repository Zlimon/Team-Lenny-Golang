package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

var (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "17"
	CONN_TYPE = "tcp"
)

func main() {
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		os.Exit(1)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Tell the server something: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
		conn.Close()
	}
}
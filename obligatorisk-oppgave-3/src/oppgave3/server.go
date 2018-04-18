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
	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer ln.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		} else {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message Received:", string(message))

			answer := "Connection established! Will now close..."
			conn.Write([]byte(answer + "\n"))
			fmt.Println("Connection closed.")
		}
		conn.Close()
	}
}
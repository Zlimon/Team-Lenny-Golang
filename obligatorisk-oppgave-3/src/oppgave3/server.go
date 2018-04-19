package main

import (
	"net"
	"fmt"
	"os"
)

var (
	connHost = "127.0.0.1"
	connPort = "17"
	connType = "tcp"
)

func main() {
	ln, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer ln.Close()
	fmt.Println("Listening on ",connHost,":",connPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		} else {
			fmt.Println("Recieved a connection!")
			fmt.Println("Returning a daily quote...")
			conn.Write([]byte("'Quote of the day!' Will now close...\n"))
			fmt.Println("Connection closed.")
		}
		conn.Close()
	}
}
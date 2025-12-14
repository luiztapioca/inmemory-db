package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("TCP server listening on 6379...")

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			reader := bufio.NewReader(c)
			for {
				message, err := reader.ReadString('\n')

				if err != nil {
					if err == io.EOF {
						fmt.Println("Connection ended.")
						break
					}
					fmt.Println("Error handling message: ", err)
					break
				}
				if strings.Trim(message, "\t\n") == "PING" {
					fmt.Println("PONG")
				}
				_, err = conn.Write([]byte(message))
				if err != nil {
					fmt.Println("Error sending response: ", err)
					break
				}
			}
		}(conn)
	}
}

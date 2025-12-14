package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}
	conn, err := ln.Accept()

	if err != nil {
		fmt.Println("Error connecting: ", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading: ", err)
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}

}

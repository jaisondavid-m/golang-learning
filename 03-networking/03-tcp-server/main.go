package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp",":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("TCP server running on port 9000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:",err)
			continue
		}
		fmt.Println("Client Connected:",conn.RemoteAddr())
		conn.Write([]byte("Hello From Go TCP Server!\n"))
		conn.Close()
	}
}
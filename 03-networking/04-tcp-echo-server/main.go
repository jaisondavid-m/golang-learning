package main

import (
	"fmt"
	"net"
)

func handleconnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte,1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Client Disconnected:",err)
			return
		}
		message := string(buffer[:n])
		fmt.Println("Recieved:",message)

		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Write Error:",err)
			return
		}
	}
}

func main() {

	listener, err := net.Listen("tcp",":9000")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("TCP Echo Server is running on :9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept Error:",err)
			continue
		}
		go handleconnection(conn)
	}

}
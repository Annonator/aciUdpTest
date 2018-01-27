package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	message := "Server: Connection established!"
	_, err := conn.WriteToUDP([]byte(message), addr)

	if err != nil {
		fmt.Printf("Could not send Response %v", err)
	}
}

func main() {
	fmt.Println("Server Started!")

	buffer := make([]byte, 2048)

	address := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}

	connection, err := net.ListenUDP("udp", &address)

	if err != nil {
		fmt.Printf("Server: Cant bind port %v\n", err)
	}

	for {
		_, remoteAddress, err := connection.ReadFromUDP(buffer)
		fmt.Printf("Server: Read Message from %v %s \n", remoteAddress, buffer)

		if err != nil {
			fmt.Printf("Server: New Error %v \n", err)
			continue
		}

		go sendResponse(connection, remoteAddress)
	}
}

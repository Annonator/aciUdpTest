package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	buffer := make([]byte, 2048)

	connection, err := net.Dial("udp", "127.0.0.1:1234")

	if err != nil {
		fmt.Printf("Client: Cant reach %v", err)
		return
	}

	fmt.Fprintf(connection, "UDP Connection established!")

	_, err = bufio.NewReader(connection).Read(buffer)

	if err == nil {
		fmt.Printf("Message: %s \n", buffer)
	} else {
		fmt.Printf("Error %v \n", err)
	}

	connection.Close()
}

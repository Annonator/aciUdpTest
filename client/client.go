package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	buffer := make([]byte, 2048)

	connection, err := net.Dial("udp", "0.0.0.0:1234")

	if err != nil {
		fmt.Printf("Client: Cant reach %v", err)
		return
	}

	fmt.Println("starting.....")

	fmt.Fprintf(connection, "UDP Connection established!")

	fmt.Println("..")

	_, err = bufio.NewReader(connection).Read(buffer)

	fmt.Println("....")

	if err == nil {
		fmt.Printf("Message: %s \n", buffer)
	} else {
		fmt.Printf("Error %v \n", err)
	}

	connection.Close()
}

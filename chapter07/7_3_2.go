package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	listner, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddress, err := listner.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %v\n", string(buffer[:length]))
	}
}

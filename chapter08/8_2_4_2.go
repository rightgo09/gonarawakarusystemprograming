package main

import (
	"net"
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	os.Remove(clientPath)
	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}
	// 送信先のアドレス
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	if err != nil {
		panic(err)
	}
	var serverAddr net.Addr = unixServerAddr
	defer conn.Close()

	fmt.Println("Sending to server")
	_, err = conn.WriteTo([]byte("Hello from client"), serverAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))
}

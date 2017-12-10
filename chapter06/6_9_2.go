package main

import (
	"net"
	"net/http"
	"fmt"
	"bufio"
	"net/http/httputil"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	var err error
	requests := make(chan *http.Request, len(sendMessages))

	conn, err = net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access: %d\n", current)
	defer conn.Close()

	// リクエストだけ先に送る
	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages) - 1
		request, err := http.NewRequest("GET", "http://localhost:8080/?message="+sendMessages[i], nil)
		if err != nil {
			panic(err)
		}
		if lastMessage {
			fmt.Println("connection close")
			//request.Header.Add("Connection", "close")
		} else {
			fmt.Println("connection keep-alive")
			request.Header.Add("Connection", "keep-alive")
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println("send:", sendMessages[i])
		requests <- request
	}
	close(requests)

	// レスポンスをまとめて受信
	reader := bufio.NewReader(conn)
	for request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		if current == len(sendMessages) {
			break
		}
	}
}

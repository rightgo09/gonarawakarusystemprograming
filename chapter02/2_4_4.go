package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	main3()
}
func main1() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}

func main2() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	req.Write(conn)
	io.Copy(os.Stdout, conn)
}

func main3handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWrite example")
}

func main3() {
	http.HandleFunc("/", main3handler)
	http.ListenAndServe(":8080", nil)
}

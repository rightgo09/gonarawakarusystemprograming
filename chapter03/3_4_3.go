package main

import (
	"io"
	"net"
	"os"
	"net/http"
	"bufio"
	"fmt"
)

func main() {
	main3_4_3_2()
}

func main3_4_3_1() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

func main3_4_3_2() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダーを表示してみる
	fmt.Println(res.Header)
	// ボディを表示してみる。最後にはClose()すること
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

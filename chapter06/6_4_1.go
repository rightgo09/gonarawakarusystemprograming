package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 一度で終了しないためにAccept()を何度も繰り返し呼ぶ
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		// 1リクエスト処理中に他のリクエストのAccept()が行えるように
		// goroutineを使って非同期にレスポンスを処理する
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))
			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello, World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}

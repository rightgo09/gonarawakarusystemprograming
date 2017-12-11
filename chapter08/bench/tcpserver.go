package bench

import (
	"bufio"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func TCPServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			// リクエストを読み込む
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			_, err = httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
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

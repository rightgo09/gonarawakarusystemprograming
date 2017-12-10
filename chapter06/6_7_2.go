package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
	"bytes"
	"compress/gzip"
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
		go processSession(conn)
	}
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// Accept後のソケットで何度も応答を返すためにループ
	for {
		fmt.Println("loop!")
		// タイムアウトを設定
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		// リクエストを読み込む
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			// タイムアウトもしくはソケットクローズ時は終了
			// それ以外はエラーにする
			neterr, ok := err.(net.Error) // ダウンキャスト
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}

		// リクエストを表示
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		// レスポンスを書き込む
		// HTTP/1.1 かつ、ContentLengthの設定が必要
		response := http.Response{
			StatusCode:    200,
			ProtoMajor:    1,
			ProtoMinor:    1,
			Header: make(http.Header),
		}
		content := "Hello World\n"
		if isGzipAceptable(request) {
			// コンテンツをgzip化して転送
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			response.Body = ioutil.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}
		response.Write(conn)
	}
}

func isGzipAceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// 青空文庫: 「ごんぎつね」より
// http://www.aozora.gr.jp/cards/000121/card628.html#download
var contents = []string{
	"これは、私《わたし》が小さいときに、村の茂平《もへい》というおじいさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山《なかやま》というところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐《ぎつね》」という狐がいました。",
	"ごんは、一人《ひとり》ぼっちの小狐で、しだ［＃「しだ」に傍点］の一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

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
		go processSession6_8_1(conn)
	}
}

func processSession6_8_1(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// Accept後のソケットで何度も応答を返すためにループ
	for {
		// リクエストを読み込む
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
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
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprint(conn, "0\r\n\r\n")
	}
}

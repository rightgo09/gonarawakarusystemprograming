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
		go processSession6_9_1(conn)
	}
}

func processSession6_9_1(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	// セッション内のリクエストを順に処理するためのチャネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	// レスポンスを直列化してソケットに書き出す専用のgoroutine
	go writeToConn(sessionResponses, conn)

	reader := bufio.NewReader(conn)

	// Accept後のソケットで何度も応答を返すためにループ
	for {
		fmt.Println("loop!")

		// レスポンスを受け取ってセッションのキューに入れる

		// タイムアウトを設定
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		// リクエストを読み込む
		request, err := http.ReadRequest(reader)
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

		sessionResopnse := make(chan *http.Response)
		sessionResponses <- sessionResopnse

		// 非同期でレスポンスを実行
		go handleRequest(request, sessionResopnse)
	}
}

func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for sessionResponse := range sessionResponses {
		// 選択された仕事が終わるまで待つ
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	// レスポンスを書き込む
	// セッションを維持するためにKeep-Aliveでないといけない
	content := "Hello World\n"
	response := &http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 1,
		ContentLength: int64(len(content)),
		Body: ioutil.NopCloser(strings.NewReader(content)),
	}
	// 処理が終わったらチャネルに書き込み、
	// ブロックされていたwriteToConnの処理を再始動する
	resultReceiver <- response
}
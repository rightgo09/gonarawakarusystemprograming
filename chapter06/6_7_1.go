package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	// リトライ用にループで全体を囲う
	for {
		var err error
		// まだコネクションを張っていない / エラーでリトライ
		if conn == nil {
			// Dial から行って conn を初期化
			conn, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest("POST", "http://localhst:8080", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		request.Header.Set("Accept-Encoding", "gzip")
		request.Write(conn)
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		// 結果を表示
		//dump, err := httputil.DumpResponse(response, true)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(string(dump))
		defer response.Body.Close()
		if response.Header.Get("Content-Encoding") == "gzip" {
			fmt.Println("gzip")
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
		} else {
			fmt.Println("plain")
			io.Copy(os.Stdout, response.Body)
		}

		// 全部送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}
		//fmt.Println("4 seconds")
		//time.Sleep(4 * time.Second)
	}
	conn.Close()
}

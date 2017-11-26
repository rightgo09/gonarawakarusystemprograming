package main

import (
	"encoding/csv"
	"os"
	"net/http"
	"encoding/json"
	"io"
	"compress/gzip"
)

func main() {
	main2_4_8_3()
}

func main2_4_8_2() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"foo", "bar", "baz"})
	writer.Write([]string{"hoge", "fuga", "piyo"})
	writer.Flush()
}

func main2_4_8_3_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	// ここにコードを書く
	gziper := gzip.NewWriter(w)
	writer := io.MultiWriter(gziper, os.Stdout)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "	")
	if err := encoder.Encode(source); err != nil {
		panic(err)
	}
	if err := gziper.Flush(); err != nil {
		panic(err)
	}
	gziper.Close()
}

func main2_4_8_3()  {
	http.HandleFunc("/", main2_4_8_3_handler)
	http.ListenAndServe(":8080", nil)
}